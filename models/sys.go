package models

type FileUploadResponse struct {
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Path     string `json:"path"`
}

type DownloadFileRequest struct {
	DirectoryName string `json:"directory_name" binding:"required"`
	FileName      string `json:"file_name" binding:"required"`
}
