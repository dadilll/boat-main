package models

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type GameState struct {
	PlayerBoards  []*GameBoard `json:"playerBoards"`
	EnemyBoards   []*GameBoard `json:"enemyBoards"`
	CurrentPlayer int          `json:"currentPlayer"`
	Players       []*Player    `json:"players"`
}

type Player struct {
	ID   int
	Conn *websocket.Conn
}

type Move struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewGameState(width, height int) *GameState {
	return &GameState{
		PlayerBoards:  []*GameBoard{NewGameBoard(width, height), NewGameBoard(width, height)},
		EnemyBoards:   []*GameBoard{NewGameBoard(width, height), NewGameBoard(width, height)},
		CurrentPlayer: 0,
		Players:       make([]*Player, 0),
	}
}

func (gs *GameState) MarshalJSON() ([]byte, error) {
	type TempGameState struct {
		PlayerBoards  []*GameBoard `json:"playerBoards"`
		EnemyBoards   []*GameBoard `json:"enemyBoards"`
		CurrentPlayer int          `json:"currentPlayer"`
		Players       []*Player    `json:"players"`
	}

	temp := TempGameState{
		PlayerBoards:  gs.PlayerBoards,
		EnemyBoards:   gs.EnemyBoards,
		CurrentPlayer: gs.CurrentPlayer,
		Players:       gs.Players,
	}

	return json.Marshal(temp)
}
