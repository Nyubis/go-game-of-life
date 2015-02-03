package gameoflife

type Board struct {
	w, h  int
	cells [][]bool
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
