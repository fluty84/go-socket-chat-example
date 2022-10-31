package controllers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"socketChatExample/server/dto"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan dto.Message)

var upgrader = websocket.Upgrader{}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true
	for {
		var msg dto.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

// Lee desde el canal y transmite a todos los clientes desde su respectivo websocket
func HandleMessages() {

	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			//log.Println("msg: ", msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
