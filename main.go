package main

import (
	"log"
	"marketplace/database"
	"marketplace/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUserByID)

	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)

	r.GET("/products", handlers.GetProducts)
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products/:id", handlers.GetProductByID)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	r.POST("/products/:id/reviews", handlers.CreateReview)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
