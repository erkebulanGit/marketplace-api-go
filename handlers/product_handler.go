package handlers

import (
	"marketplace/database"
	"marketplace/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	query := database.DB.Preload("User").Preload("Category")

	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if userID := c.Query("user_id"); userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	var total int64
	query.Model(&models.Product{}).Count(&total)
	query.Limit(limit).Offset((page - 1) * limit).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"data": products, "total": total, "page": page, "limit": limit,
	})
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if product.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if product.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than 0"})
		return
	}
	if product.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID is required"})
		return
	}
	if product.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CategoryID is required"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, product.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var category models.Category
	if err := database.DB.First(&category, product.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	database.DB.Create(&product)
	database.DB.Preload("User").Preload("Category").First(&product, product.ID)
	c.JSON(http.StatusCreated, product)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.Preload("User").Preload("Category").Preload("Reviews.User").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if input.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if input.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than 0"})
		return
	}

	database.DB.Model(&product).Updates(models.Product{
		Title: input.Title, Description: input.Description,
		Price: input.Price, CategoryID: input.CategoryID,
	})

	database.DB.Preload("User").Preload("Category").First(&product, product.ID)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	database.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
