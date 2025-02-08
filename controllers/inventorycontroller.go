package controllers

import (
	"assigment/config"
	"assigment/models"
	"assigment/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetInventoryHandler retrieves the inventory details for a specific product.
func GetInventoryHandler(c *gin.Context) {
	// Get product_id from URL parameter
	productIDStr := c.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32) // convert string to uint
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	// Retrieve inventory for the given product ID
	inventory, err := models.GetInventoryByProductID(config.DB, uint(productID))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Inventory not found")
		return
	}

	// Return success response with inventory details
	utils.SendSuccessResponse(c, http.StatusOK, "Inventory details", inventory)
}

// UpdateInventoryHandler updates the inventory for a specific product.
func UpdateInventoryHandler(c *gin.Context) {
	var inventory models.Inventory

	// Bind incoming JSON request body to the inventory struct
	if err := c.ShouldBindJSON(&inventory); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Get product_id from URL parameter
	productIDStr := c.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32) // convert string to uint
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	// Check if the product exists before updating inventory
	product, err := models.GetProductByID(config.DB, uint(productID))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	// Ensure inventory changes are only made for the valid product
	inventory.ProductID = product.ID

	// Update inventory for the given product
	if err := models.UpdateInventory(config.DB, uint(productID), &inventory); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update inventory")
		return
	}

	// Return success response with updated inventory
	utils.SendSuccessResponse(c, http.StatusOK, "Inventory updated successfully", inventory)
}
