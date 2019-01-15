package main

// http://codingdojo.org/kata/GameOfLife/

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	DeadCell Cell = false
	LiveCell Cell = true
)

type Cell bool

func (c Cell) String() string {
	if c {
		return "*"
	}

	return "."
}

type Generation struct {
	Num   int
	Cells [][]Cell
}

func newGeneration(size int) Generation {
	cells := make([][]Cell, size, size)

	for i := range cells {
		cells[i] = make([]Cell, size, size)

		for j := range cells[i] {
			cells[i][j] = DeadCell
		}
	}

	return Generation{Num: 1, Cells: cells}
}

func randGeneration(size int) Generation {
	cells := make([][]Cell, size, size)

	for i := range cells {
		cells[i] = make([]Cell, size, size)

		for j := range cells[i] {
			cells[i][j] = rand.Intn(2) == 0
		}
	}

	return Generation{Num: 1, Cells: cells}
}

func (gen Generation) String() string {
	str := fmt.Sprintf("%d. Generation\n", gen.Num)

	for i := range gen.Cells {
		for _, cell := range gen.Cells[i] {
			str += cell.String()
		}

		str += "\n"
	}

	return str
}

func (gen Generation) NeighboursOfCell(x, y int) (n int) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {

			if i == 0 && j == 0 {
				continue // skip self
			}

			if x+i < 0 || y+j < 0 {
				continue // skip index out of range
			}

			if x+i >= len(gen.Cells) || y+j >= len(gen.Cells[0]) {
				continue // skip index out of range
			}

			if gen.Cells[x+i][y+j] == LiveCell {
				n++
			}
		}
	}

	return
}

func (gen Generation) NextGeneration() Generation {
	nextGen := newGeneration(len(gen.Cells))
	nextGen.Num = gen.Num + 1

	for i := range gen.Cells {
		for j, cell := range gen.Cells[i] {

			n := gen.NeighboursOfCell(i, j)
			switch {
			// Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
			case n < 2:
				nextGen.Cells[i][j] = DeadCell
			// Any live cell with more than three live neighbours dies, as if by overcrowding.
			case n > 3:
				nextGen.Cells[i][j] = DeadCell
			// Any live cell with two or three live neighbours lives on to the next generation.
			case n == 2 && cell == LiveCell:
				nextGen.Cells[i][j] = LiveCell
			// Any dead cell with exactly three live neighbours becomes a live cell.
			case n == 3:
				nextGen.Cells[i][j] = LiveCell
			}
		}
	}

	return nextGen
}

func main() {
	gen := randGeneration(8)
	for {
		fmt.Println(gen)

		nextGen := gen.NextGeneration()

		if reflect.DeepEqual(gen.Cells, nextGen.Cells) {
			fmt.Println(nextGen)
			break
		}

		gen = nextGen
	}
}
