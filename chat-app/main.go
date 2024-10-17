package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var bannedWords = []string{"fuck", "shit", "dick"} // Добавьте сюда запрещенные слова

func containsBannedWords(message string) bool {
	for _, word := range bannedWords {
		if strings.Contains(strings.ToLower(message), word) {
			return true
		}
	}
	return false
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v\n", err)
		return
	}
	defer ws.Close()

	clients[ws] = true
	log.Printf("New client connected: %v\n", ws.RemoteAddr())

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading JSON: %v\n", err)
			delete(clients, ws)
			break
		}

		if containsBannedWords(msg.Message) {
			log.Printf("Message from %s contains banned words and will not be broadcasted.\n", msg.Username)
			continue
		}

		log.Printf("Received message from %s: %s\n", msg.Username, msg.Message)
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing JSON to client: %v\n", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v\n", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("HTTP server started on :8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("ListenAndServe: %v\n", err)
	}
}
