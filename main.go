package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	
	// We don't really care for origin right now :)
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		// Upgrades http to ws.
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Print("Upgrade failed: ", err)
			return
		}

		defer conn.Close()

		// Continuosly read and write message
		for {
			// Reads message from stream
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read failed:", err)
				break
			}
				
			// Prints to console
			log.Println("Received: ", string(message))

			
			// Echoes message back to stream
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("Write failed:", err)
				break
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Server file for testing.
		http.ServeFile(w, r, "index.html")
	})

	log.Println("Running serve on port 8080")
	http.ListenAndServe(":8080", nil)
}
