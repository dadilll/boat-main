package domain

import (
	"encoding/json"
	"log"

	"github.com/Dadil/boat/backend/Gamelogic/models"
	"github.com/gorilla/websocket"
)

// CheckHit проверяет попадание в корабль
func CheckHit(board *models.GameBoard, x, y int) bool {
	if x < 0 || y < 0 || x >= board.Width || y >= board.Height {
		// Проверка, что координаты находятся в пределах игрового поля
		return false
	}
	return board.Cells[y][x] == models.CellStateShip
}

func ProcessMove(gameState *models.GameState, playerID int, move models.Move) {
	enemyPlayerID := (playerID + 1) % 2

	enemyBoard := gameState.PlayerBoards[enemyPlayerID]
	playerBoard := gameState.EnemyBoards[playerID]

	// Обработка выстрела
	if CheckHit(enemyBoard, move.X, move.Y) {
		// Если попадание
		enemyBoard.Cells[move.Y][move.X] = models.CellStateHit
		playerBoard.Cells[move.Y][move.X] = models.CellStateHit
	} else {
		// Если промах
		enemyBoard.Cells[move.Y][move.X] = models.CellStateMiss
		playerBoard.Cells[move.Y][move.X] = models.CellStateMiss
	}

	log.Printf("Processed move by player %d: (%d, %d), current game state: %+v", playerID, move.X, move.Y, gameState)

	for _, player := range gameState.Players {
		if err := SendGameState(player.Conn, gameState); err != nil {
			log.Println("Failed to send game state:", err)
		}
	}
}

func SendGameState(conn *websocket.Conn, gameState *models.GameState) error {
	gameStateJSON, err := json.Marshal(gameState)
	if err != nil {
		return err
	}

	// Логирование данных перед отправкой
	log.Printf("Sending game state: %s", gameStateJSON)

	if err := conn.WriteMessage(websocket.TextMessage, gameStateJSON); err != nil {
		return err
	}

	return nil
}
