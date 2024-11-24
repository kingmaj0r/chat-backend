package src

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	clients    = make(map[*websocket.Conn]string)
	clientsMux sync.Mutex
	usernames  = make(map[string]bool)
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	var initialMessage Message
	err = conn.ReadJSON(&initialMessage)
	if err != nil {
		log.Println("Initial connection error:", err)
		return
	}

	clientsMux.Lock()
	if usernames[initialMessage.Sender] {
		conn.WriteJSON(Message{Sender: "Server", Text: "Username already taken"})
		clientsMux.Unlock()
		return
	}

	clients[conn] = initialMessage.Sender
	usernames[initialMessage.Sender] = true
	clientsMux.Unlock()

	defer func() {
		clientsMux.Lock()
		delete(usernames, initialMessage.Sender)
		delete(clients, conn)
		clientsMux.Unlock()
	}()

	broadcast(fmt.Sprintf("%s has joined the chat", initialMessage.Sender), "Server", conn)

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		broadcast(msg.Text, msg.Sender, conn)
	}
}

func broadcast(messageText string, sender string, ignoreConn *websocket.Conn) {
	clientsMux.Lock()
	defer clientsMux.Unlock()

	for client, username := range clients {
		if client != ignoreConn {
			err := client.WriteJSON(Message{Sender: sender, Text: messageText})
			if err != nil {
				log.Printf("Error sending message: %v", err)
				client.Close()
				delete(usernames, username)
				delete(clients, client)
			}
		} else {
			err := client.WriteJSON(Message{Sender: sender, Text: messageText})
			if err != nil {
				log.Printf("Error sending message to sender: %v", err)
				client.Close()
				delete(usernames, username)
				delete(clients, client)
			}
		}
	}
}
