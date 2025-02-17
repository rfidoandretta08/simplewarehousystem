package controllers

import (
	"assigment/models"
	"assigment/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	idStr := c.Param("id")

	_, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}
	const uploadDir = "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(400, gin.H{"error": "Gagal membuat directory"})
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{"error": "File tidak boleh kosong"})
		return
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpeg" && ext != ".jpg" {
		c.JSON(400, gin.H{"error": "Hanya file JPEG (.jpeg atau .jpg) yang diperbolehkan"})
		return
	}

	if file.Size > 1*1024*1024 {
		c.JSON(400, gin.H{"error": "Ukuran file melebihi batas (maksimum 1 MB)"})
		return
	}

	filename := filepath.Base(file.Filename)

	//"example_dir3/belajarGolang-Edit.txt", (contoh)
	PathFolder := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, PathFolder); err != nil {
		c.JSON(400, gin.H{"error": "Error mengupload file"})
		return
	}

	response := models.FileUploadResponse{
		Message:  "File berhasil diupload!",
		Filename: filename,
		Path:     PathFolder,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})

}

func DownloadFile(c *gin.Context) {
	idStr := c.Param("id")

	_, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	fileName := c.Query("file_name")
	dirName := c.Query("directory_name")

	if fileName == "" || dirName == "" {
		c.JSON(400, gin.H{"error": "Filename dan directory name tidak boleh kosong!"})
		return
	}

	//"example_dir3/belajarGolang-Edit.txt", (contoh)
	filePath := filepath.Join(dirName, fileName)

	absPath, err := filepath.Abs(filePath)

	if err != nil {
		c.JSON(400, gin.H{"error": "Folder tidak ditemukan!"})
		return
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		c.JSON(400, gin.H{"error": "File tidak ditemukan!"})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Type", "application,octet-stream")
	// bikin channel
	done := make(chan bool)
	errChan := make(chan error)

	// bikin fungsi async

	go func() {
		file, err := os.Open(absPath)
		if err != nil {
			c.JSON(400, gin.H{"error": "Gagal membuka file!"})
			return
		}

		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			c.JSON(400, gin.H{"error": "Gagal membaca file"})
			return
		}

		c.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

		buffer := make([]byte, 32*1024) // Read 32KB at a time
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				errChan <- fmt.Errorf("error reading file: %v", err)
				return
			}

			if _, err := c.Writer.Write(buffer[:n]); err != nil {
				errChan <- fmt.Errorf("error writing to response: %v", err)
				return
			}
			c.Writer.Flush() // Flush the response writer after each chunk
		}
	}()

	select {
	case <-done:
		return
	case err := <-errChan:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	case <-time.After(5 * time.Minute): // Timeout after 5 minutes
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Download timeout"})
		return
	}
}
