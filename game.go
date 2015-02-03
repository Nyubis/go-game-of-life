package gameoflife

type Board struct {
	w, h  int
	cells [][]bool //A collection of rows, which contain cells
}

func MakeBoard(width int, height int) *Board {
	cells := make([][]bool, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]bool, width)
	}
	return &Board{
		w:     width,
		h:     height,
		cells: cells,
	}
}

func (board *Board) GetAt(x int, y int) bool {
	if (x < 0 || y < 0 || x >= board.w || y >= board.h) {
		return false
	}
	return board.cells[y][x]
}

func (board *Board) SetAt(x int, y int, val bool) {
	if (x < 0 || y < 0 || x >= board.w || y >= board.h) {
		return // TODO: error reporting
	}
	board.cells[y][x] = val
}
