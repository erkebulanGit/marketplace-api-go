package main

import (
	"log"
	"marketplace/database"
	"marketplace/handlers"
	"marketplace/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProductByID)
	r.GET("/categories", handlers.GetCategories)

	protected := r.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/auth/me", handlers.Me)
		protected.GET("/users/:id", handlers.GetUserByID)
		protected.POST("/categories", handlers.CreateCategory)
		protected.POST("/products", handlers.CreateProduct)
		protected.PUT("/products/:id", handlers.UpdateProduct)
		protected.DELETE("/products/:id", handlers.DeleteProduct)
		protected.POST("/products/:id/reviews", handlers.CreateReview)
	}

	log.Println("Server running on :8080")
	r.Run(":8080")
}
