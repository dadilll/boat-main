package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dadil/boat/backend/Gamelogic/domain"
	"github.com/Dadil/boat/backend/Gamelogic/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // УБРАТЬ ПОТОМ!!!!!!!!!!
	},
}

func HandleConnections(room *Room, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	player := &models.Player{Conn: ws}
	room.AddPlayer(player)
	defer room.RemovePlayer(player)

	player.ID = len(room.GameState.Players) - 1

	// Send initial game state and player ID to the new player
	initMsg := map[string]interface{}{
		"type":     "init",
		"playerID": player.ID,
	}
	if err := ws.WriteJSON(initMsg); err != nil {
		log.Println("Error sending init message:", err)
		return
	}
	if err := domain.SendGameState(ws, room.GameState); err != nil {
		log.Println("Error sending initial game state:", err)
		return
	}

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var move models.Move
		if err := json.Unmarshal(msg, &move); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		room.mu.Lock()
		domain.ProcessMove(room.GameState, player.ID, move)
		room.mu.Unlock()
	}
}
