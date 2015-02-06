package gameoflife

type Board struct {
	w, h   int
	cells  [][]bool // A collection of rows, which contain cells
	shadow [][]bool // Start as copy of cells, used as temp for calculating next step
}

func MakeBoard(width int, height int) *Board {
	cells := make([][]bool, height)
	shadow := make([][]bool, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]bool, width)
		shadow[i] = make([]bool, width)
	}
	return &Board{
		w:      width,
		h:      height,
		cells:  cells,
		shadow: shadow,
	}
}

func (board *Board) GetAt(x int, y int) bool {
	if x < 0 || y < 0 || x >= board.w || y >= board.h {
		return false
	}
	return board.cells[y][x]
}

func (board *Board) GetCells() [][]bool {
	return board.cells
}

func (board *Board) SetAt(x int, y int, val bool) {
	if x < 0 || y < 0 || x >= board.w || y >= board.h {
		return // TODO: error reporting
	}
	board.cells[y][x] = val
}

func (board *Board) Step() {
	for i, row := range board.cells {
		for j, cell := range row {
			n := board.countNeighbours(i,j)
			board.shadow[i][j] = checkLive(n, cell)
		}
	}
	board.shadow, board.cells = board.cells, board.shadow
}

func (board *Board) countNeighbours(row int, col int) int {
	var count int
	var leftbound, rightbound, upbound, lowbound int // In case the cell is at the edge of the board
	leftbound, rightbound = calcBounds(col, board.w)
	upbound, lowbound = calcBounds(row, board.h)

	for i := upbound; i <= lowbound; i++ {
		for j := leftbound; j <= rightbound; j++ {
			if board.cells[i][j] && (row != i || col != j){
				count++
			}
		}
	}

	return count
}

func calcBounds(val int, max int) (int, int) {
	var minbound, maxbound int
	if val == 0 {
		minbound = 0
	} else {
		minbound = val - 1
	}
	if val == max - 1 {
		maxbound = val
	} else {
		maxbound = val + 1
	}
	return minbound, maxbound
}

func checkLive(neighs int, isAlive bool) bool {
	if neighs <= 1 || neighs >= 4 {
		return false // Die of isolation or overpopulation
	} else {
		return isAlive || neighs == 3 // Survive or be reborn
	}
}
