package main

import (
	"log"
	"net/http"
	"socketChatExample/server/controllers"
)

func main() {
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", controllers.HandleConnections)

	go controllers.HandleMessages()

	log.Println("http server started on :8080")
	// uncomment to deploy on vercel err := http.ListenAndServeTLS(":8080", "ca-cert.pem", "ca-key.pem", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
