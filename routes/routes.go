package routes

import (
	"assigment/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rute untuk Product
	router.POST("/products", controllers.CreateProductHandler)                           // Menambahkan produk
	router.GET("/products/:id", controllers.GetProductHandler)                           // Melihat produk berdasarkan ID
	router.GET("/products/category/:category", controllers.GetProductsByCategoryHandler) // Melihat produk berdasarkan kategori
	router.PUT("/products/:id", controllers.UpdateProductHandler)                        // Memperbarui produk
	router.DELETE("/products/:id", controllers.DeleteProductHandler)                     // Menghapus produk

	// Rute untuk Inventory
	router.GET("/inventories/:product_id", controllers.GetInventoryHandler)    // Melihat stok produk
	router.PUT("/inventories/:product_id", controllers.UpdateInventoryHandler) // Memperbarui stok produk

	// Rute untuk orders
	router.POST("/orders/", controllers.CreateOrderHandler) // Membuat pesanan baru
	router.GET("/orders/:id", controllers.GetOrderHandler)  // Melihat pesanan berdasarkan ID

	//rute upload dan download
	router.POST("products/upload/:id", controllers.UploadFile)
	router.GET("/products/file/download/:id", controllers.DownloadFile)

}
