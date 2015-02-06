package main

import(
	"bytes"
	"fmt"
	"time"
	"github.com/nyubis/game-of-life/gameoflife"
)

func main() {
	b := gameoflife.MakeBoard(10,10)
	// .x.
	// ..x
	// xxx
	b.SetAt(1,0,true)
	b.SetAt(2,1,true)
	b.SetAt(0,2,true)
	b.SetAt(1,2,true)
	b.SetAt(2,2,true)

	for {
		fmt.Println("\033[1;1H\033[2J") // Black magic to clear the screen
		fmt.Println(render(b.GetCells(), '░', '▓'))
		b.Step()
		time.Sleep(250 * time.Millisecond)
	}
}

func render(cells [][]bool, dead rune, alive rune) string {
	var output bytes.Buffer
	for _, row := range cells {
		for _, cell := range row {
			if cell {
				output.WriteRune(alive)
			} else {
				output.WriteRune(dead)
			}
		}
		output.WriteRune('\n')
	}
	return output.String()
}
