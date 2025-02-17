package controllers

import (
	"assigment/config"
	"assigment/models"
	"assigment/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//menampilkan data dari inventory by product_id
func GetInventoryHandler(c *gin.Context) {
	// Get product_id from URL parameter
	productIDStr := c.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	inventory, err := models.GetInventoryByProductID(config.DB, uint(productID))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Inventory not found")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Inventory details", inventory)
}

//mengupdate data stock , lokasi by product_id
func UpdateInventoryHandler(c *gin.Context) {
	var inventory models.Inventory

	if err := c.ShouldBindJSON(&inventory); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	productIDStr := c.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32) // convert string to uint
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := models.GetProductByID(config.DB, uint(productID))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	inventory.ProductID = product.ID

	if err := models.UpdateInventory(config.DB, uint(productID), &inventory); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update inventory")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Inventory updated successfully", inventory)
}
