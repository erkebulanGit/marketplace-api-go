package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Notification struct {
	ID           int       `json:"id"`
	UserEmail    string    `json:"user_email"`
	ProductTitle string    `json:"product_title"`
	Message      string    `json:"message"`
	CreatedAt    time.Time `json:"created_at"`
}

var notifications []Notification
var nextID = 1

func main() {
	r := gin.Default()

	r.POST("/notify", func(c *gin.Context) {
		var notif Notification

		if err := c.ShouldBindJSON(&notif); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		notif.ID = nextID
		notif.CreatedAt = time.Now()
		nextID++
		notifications = append(notifications, notif)

		log.Printf("Уведомление #%d отправлено на %s: '%s'",
			notif.ID, notif.UserEmail, notif.ProductTitle)

		c.JSON(http.StatusOK, gin.H{
			"message": "Notification sent successfully",
			"to":      notif.UserEmail,
			"about":   notif.ProductTitle,
		})
	})

	r.GET("/notifications", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"total": len(notifications),
			"data":  notifications,
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Notification Service running on :8081")
	r.Run(":8081")
}
