// controllers/order_controller.go

package controllers

import (
	"assigment/config"
	"assigment/models"
	"assigment/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// membuat orderan baru
func CreateOrderHandler(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("Received product_id: %d", order.ProductID)

	if order.ProductID == 0 {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := models.GetProductByID(config.DB, order.ProductID)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	log.Printf("Found product: %v", product)

	if err := models.CreateOrder(config.DB, &order); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create order")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Order created successfully", order)
}

// menampilkan order by product_id
func GetOrderHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	order, err := models.GetOrderByID(config.DB, uint(id))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Order not found")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Order details", order)
}
