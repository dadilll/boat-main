package domain

import (
	"math/rand"
	"time"

	"github.com/Dadil/boat/backend/Gamelogic/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func InitializeGameBoard() *models.GameBoard {
	board := &models.GameBoard{
		Width:  10,
		Height: 10,
		Cells:  make([][]models.CellState, 10),
	}
	for i := range board.Cells {
		board.Cells[i] = make([]models.CellState, 10)
	}
	return board
}

func PlaceShips(board *models.GameBoard, ships []*models.Ship) error {
	for _, ship := range ships {
		placed := false
		for !placed {
			x := rand.Intn(board.Width)
			y := rand.Intn(board.Height)
			if isValidPlacement(board, ship, models.Coordinate{X: x, Y: y}) {
				placeShip(board, ship, []models.Coordinate{{X: x, Y: y}})
				placed = true
			}
		}
	}
	return nil
}

func isValidPlacement(board *models.GameBoard, ship *models.Ship, coord models.Coordinate) bool {
	if coord.X+ship.Width > board.Width || coord.Y+ship.Height > board.Height {
		return false
	}
	for i := 0; i < ship.Width; i++ {
		for j := 0; j < ship.Height; j++ {
			if board.Cells[coord.Y+j][coord.X+i] != models.CellStateEmpty {
				return false
			}
		}
	}
	return true
}

func placeShip(board *models.GameBoard, ship *models.Ship, coords []models.Coordinate) {
	for _, coord := range coords {
		for i := 0; i < ship.Width; i++ {
			for j := 0; j < ship.Height; j++ {
				board.Cells[coord.Y+j][coord.X+i] = models.CellStateShip
			}
		}
	}
}
