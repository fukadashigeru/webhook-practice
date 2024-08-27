package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", handleWebhook)
	http.HandleFunc("/webhook/register", handleRegister)
	fmt.Println("Webhook receiver is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("Received webhook: %s\n", string(body))
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		challenge := r.URL.Query().Get("hub.challenge")
		w.Write([]byte(challenge))
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
