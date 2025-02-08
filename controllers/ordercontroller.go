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

func CreateOrderHandler(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		// Mengirimkan response error jika binding JSON gagal
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Debugging: Log untuk memeriksa apakah product_id sudah terbind dengan benar
	log.Printf("Received product_id: %d", order.ProductID)

	// Validasi product_id yang diterima
	if order.ProductID == 0 {
		// Mengirimkan response error jika product_id tidak valid
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	// Cek apakah produk dengan product_id ada di database
	product, err := models.GetProductByID(config.DB, order.ProductID)
	if err != nil {
		// Jika produk tidak ditemukan, kirimkan response error
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	// Log untuk memeriksa produk yang ditemukan
	log.Printf("Found product: %v", product)

	// Buat pesanan baru
	if err := models.CreateOrder(config.DB, &order); err != nil {
		// Kirimkan response error jika pembuatan pesanan gagal
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create order")
		return
	}

	// Kirimkan response sukses jika pesanan berhasil dibuat
	utils.SendSuccessResponse(c, http.StatusOK, "Order created successfully", order)
}

// GetOrderHandler untuk mengambil pesanan berdasarkan ID
func GetOrderHandler(c *gin.Context) {
	idStr := c.Param("id")                      // Ambil id sebagai string
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert id menjadi uint
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
