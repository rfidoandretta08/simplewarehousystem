package controllers

import (
	"assigment/config"
	"assigment/models"
	"assigment/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func untuk membuat produk baru
func CreateProductHandler(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateProduct(config.DB, &product); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	inventory := models.Inventory{
		ProductID: product.ID,
		Quantity:  0,
	}

	if err := models.CreateInventory(config.DB, &inventory); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create inventory")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Product created successfully with inventory", product)
}

// func untuk menampilkan product by id
func GetProductHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := models.GetProductByID(config.DB, uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Product details", product)
}

// func untuk mengupdate data dari querry product by id
func UpdateProductHandler(c *gin.Context) {
	var product models.Product
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := models.UpdateProduct(config.DB, uint(id), &product); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update product")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Product updated successfully", product)
}

// func untuk menghapus data dari querry product by id
func DeleteProductHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := models.DeleteProduct(config.DB, uint(id)); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Product deleted successfully", nil)
}

// func untuk mendapatkan data by categori
func GetProductsByCategoryHandler(c *gin.Context) {
	category := c.Param("category")
	products, err := models.GetProductsByCategory(config.DB, category)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Products not found")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Products details", products)
}
