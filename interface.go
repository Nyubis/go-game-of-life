package main

import(
	"bytes"
	"fmt"
	"time"
	"os"
	"io/ioutil"
	"strings"
	"errors"
	"github.com/nyubis/game-of-life/gameoflife"
)

func main() {
	var b *gameoflife.Board
	if len(os.Args) == 1 {
		b = example()
	} else {
		var err error
		b, err = parseFromFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

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

func example() *gameoflife.Board {
	b := gameoflife.MakeBoard(10,10)
	// .x.
	// ..x
	// xxx
	b.SetAt(1,0,true)
	b.SetAt(2,1,true)
	b.SetAt(0,2,true)
	b.SetAt(1,2,true)
	b.SetAt(2,2,true)
	return b
}

func parseFromFile (path string) (*gameoflife.Board, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) == 0 {
		return nil, errors.New("empty file")
	}
	b := gameoflife.MakeBoard(len(lines[0]), len(lines))
	for y, line := range lines {
		for x, char := range line {
			// Just 2 representations for “dead” for now
			if (char != ' ' && char != '.') {
				b.SetAt(x, y, true)
			}
		}
	}

	return b, nil
}
