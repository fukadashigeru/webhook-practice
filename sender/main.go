package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	webhookURL := "http://localhost:8080/webhook"

	// まずWebhookの登録（検証）を行う
	registerWebhook(webhookURL)

	// 定期的にWebhookを送信
	for {
		sendWebhook(webhookURL)
		time.Sleep(5 * time.Second)
	}
}

func registerWebhook(url string) {
	resp, err := http.Get(url + "/register?hub.challenge=testchallenge")
	if err != nil {
		fmt.Printf("Error registering webhook: %v\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Webhook registered with status: %s\n", resp.Status)
}

func sendWebhook(url string) {
	payload := []byte(fmt.Sprintf(`{"event": "update", "time": "%s"}`, time.Now().Format(time.RFC3339)))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Printf("Error sending webhook: %v\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Webhook sent with status: %s\n", resp.Status)
}
