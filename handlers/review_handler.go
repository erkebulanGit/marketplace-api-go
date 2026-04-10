package handlers

import (
	"marketplace/database"
	"marketplace/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if review.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID is required"})
		return
	}
	if review.Rating < 1 || review.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5"})
		return
	}

	review.ProductID = uint(productID)
	database.DB.Create(&review)
	database.DB.Preload("User").First(&review, review.ID)
	c.JSON(http.StatusCreated, review)
}
