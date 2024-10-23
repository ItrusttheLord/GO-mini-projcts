package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (for testing purposes)
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ClientManager struct {
	clients    map[*Client]bool // Track connected clients
	register   chan *Client     // Channel for registering clients
	unregister chan *Client     // Channel for unregistering clients
	broadcast  chan []byte      // Channel for broadcasting messages
}

type Client struct {
	conn *websocket.Conn // The WebSocket connection for this client
	send chan []byte     // A channel for sending messages
}

func (cm *ClientManager) start() {
	for {
		select {
		case client := <-cm.register:
			cm.clients[client] = true
		case client := <-cm.unregister:
			delete(cm.clients, client)
		case msg := <-cm.broadcast:
			for client := range cm.clients {
				select {
				case client.send <- msg: // Send the message to all clients
				default:
					close(client.send)
					delete(cm.clients, client)
				}
			}
		}
	}
}

func (client *Client) writePump() {
	defer client.conn.Close()

	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				// The channel is closed, meaning the connection should close
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// Write the message to the WebSocket
			err := client.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request, cm *ClientManager) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}

	client := &Client{
		conn: conn,
		send: make(chan []byte),
	}

	cm.register <- client

	// Start writePump for sending messages
	go client.writePump()

	// Start reading messages from the WebSocket
	go func() {
		defer func() {
			cm.unregister <- client
			client.conn.Close()
		}()

		for {
			_, msg, err := client.conn.ReadMessage()
			if err != nil {
				return
			}
			cm.broadcast <- msg
		}
	}()
}

func main() {
	cm := &ClientManager{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
	go cm.start()

	// Serve the WebSocket connection
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(w, r, cm)
	})

	// Serve the HTML file (index.html)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Serve static assets (like CSS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
