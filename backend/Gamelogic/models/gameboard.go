package models

// CellState представляет состояние клетки игрового поля
type CellState int

const (
	CellStateEmpty CellState = iota
	CellStateShip
	CellStateHit
	CellStateMiss
)

// Coordinate представляет координаты на игровом поле
type Coordinate struct {
	X int
	Y int
}

// GameBoard представляет игровое поле
type GameBoard struct {
	Width  int
	Height int
	Cells  [][]CellState
}

// NewGameBoard создает новое игровое поле заданного размера
func NewGameBoard(width, height int) *GameBoard {
	cells := make([][]CellState, height)
	for i := range cells {
		cells[i] = make([]CellState, width)
	}
	return &GameBoard{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}
