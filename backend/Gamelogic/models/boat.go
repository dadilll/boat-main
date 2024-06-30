package models

type Ship struct {
	Width  int
	Height int
	Cells  [][]CellState
}

// NewShip создает новый корабль заданного размера
func NewShip(width, height int) *Ship {
	cells := make([][]CellState, height)
	for i := range cells {
		cells[i] = make([]CellState, width)
		for j := range cells[i] {
			cells[i][j] = CellStateShip
		}
	}
	return &Ship{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

// SingleDeckShip представляет однопалубный корабль
type SingleDeckShip struct {
	Ship
}

// NewSingleDeckShip создает новый однопалубный корабль
func NewSingleDeckShip() *SingleDeckShip {
	return &SingleDeckShip{
		Ship: *NewShip(1, 1),
	}
}

// DoubleDeckShip представляет двупалубный корабль
type DoubleDeckShip struct {
	Ship
}

// NewDoubleDeckShip создает новый двупалубный корабль
func NewDoubleDeckShip() *DoubleDeckShip {
	return &DoubleDeckShip{
		Ship: *NewShip(1, 2),
	}
}

// Threedeckship представляет трехпалубный корабль
type Threedeckship struct {
	Ship
}

// NewThreedeckship создает новый трехпалубный корабль
func NewThreedeckship() *Threedeckship {
	return &Threedeckship{
		Ship: *NewShip(1, 3),
	}
}

// Fourdeckship представляет четырехпалубный корабль
type Fourdeckship struct {
	Ship
}

// NewFourdeckship создает новый четырехпалубный корабль
func NewFourdeckship() *Fourdeckship {
	return &Fourdeckship{
		Ship: *NewShip(1, 4),
	}
}
