package main

import (
	"log"
	"net/http"

	"github.com/Dadil/boat/backend/server"
)

func main() {
	room := server.NewRoom("room1")

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.HandleConnections(room, w, r)
	})

	addr := "localhost:8080"
	log.Println("http server started on", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
