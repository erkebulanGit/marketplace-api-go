package handlers

import (
	"log"

	"github.com/go-resty/resty/v2"
)

const notificationServiceURL = "http://localhost:8081"

type NotificationPayload struct {
	UserEmail    string `json:"user_email"`
	ProductTitle string `json:"product_title"`
	Message      string `json:"message"`
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
		Post(notificationServiceURL + "/notify")

	if err != nil {
		log.Println("Failed to send notification:", err)
		return
	}

	log.Println("Notification service response:", resp.Status())
}
