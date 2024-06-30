package server

import (
	"sync"

	"github.com/Dadil/boat/backend/Gamelogic/models"
)

type Room struct {
	ID        string
	Players   map[*models.Player]bool
	mu        sync.Mutex
	GameState *models.GameState
}

func NewRoom(id string) *Room {
	return &Room{
		ID:        id,
		Players:   make(map[*models.Player]bool),
		GameState: models.NewGameState(10, 10), // Используем размер 10x10
	}
}

func (r *Room) AddPlayer(player *models.Player) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Players[player] = true
	player.ID = len(r.GameState.Players)
	r.GameState.Players = append(r.GameState.Players, player)
}

func (r *Room) RemovePlayer(player *models.Player) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.Players[player]; ok {
		delete(r.Players, player)
		player.Conn.Close()

		for i, p := range r.GameState.Players {
			if p == player {
				r.GameState.Players = append(r.GameState.Players[:i], r.GameState.Players[i+1:]...)
				break
			}
		}
	}
}

func (r *Room) GetPlayerByID(id int) *models.Player {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, player := range r.GameState.Players {
		if player.ID == id {
			return player
		}
	}
	return nil
}
