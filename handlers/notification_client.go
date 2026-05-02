package handlers

import (
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

type NotificationPayload struct {
	UserEmail    string `json:"user_email"`
	ProductTitle string `json:"product_title"`
	Message      string `json:"message"`
}

func getNotificationURL() string {
	url := os.Getenv("NOTIFICATION_SERVICE_URL")
	if url == "" {
		url = "http://localhost:8081"
	}
	return url
}

func sendNotification(email string, productTitle string) {
	client := resty.New()

	payload := NotificationPayload{
		UserEmail:    email,
		ProductTitle: productTitle,
		Message:      "Ваш продукт успешно создан на маркетплейсе!",
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(getNotificationURL() + "/notify")

	if err != nil {
		log.Println("Failed to send notification:", err)
		return
	}

	log.Println("Notification response:", resp.Status())
}
