package main

import (
	"log"
	"net/http"
	"chat-backend/src"
)

func main() {
	http.HandleFunc("/ws", src.HandleConnections)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
