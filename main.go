package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cell struct {
	x int
	y int
	alive bool
	nextAliveStatus bool
}

type game struct {
	rows int
	cols int
	cellsMatrix [][]*cell
}

func (game *game) genRandomState() {
	// Generate a random alive value for every cell
	for row, _ := range game.cellsMatrix {
		for _, cell := range game.cellsMatrix[row] {			
			rand.Seed(time.Now().UnixNano())
			cell.alive = rand.Intn(7) == 1
		}
	}
}

func (game *game) getAliveCellNeighbors(cell *cell) int {
	// Find the number of alive cell neighbors of the given cell
	var aliveNeighbors int
	around := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, positions := range around {		
		x := cell.x + positions[1]
		y := cell.y + positions[0]

		// Handle cases when we hit a wall
		if x < 0 || y < 0 || x >= game.cols || y >= game.rows {
			continue
		}
		if game.cellsMatrix[y][x].alive {
			aliveNeighbors++
		}
	}

	return aliveNeighbors
}

func (game *game) nextGeneration() {
	// Calculate the next generation of the game following the 4 Conway's Game of Life rules
	for row, _ := range game.cellsMatrix {
		for _, cell := range game.cellsMatrix[row] {
			aliveNeighbors := game.getAliveCellNeighbors(cell)
			
			if (cell.alive && aliveNeighbors <= 2) {cell.nextAliveStatus = false}
			if (cell.alive && aliveNeighbors > 3) {cell.nextAliveStatus = false}
			if (cell.alive && (aliveNeighbors == 2 || aliveNeighbors == 3)) {cell.nextAliveStatus = true}
			if (!cell.alive && aliveNeighbors == 3) {cell.nextAliveStatus = true}
		}
	}

	for row, _ := range game.cellsMatrix {
		for _, cell := range game.cellsMatrix[row] {
			cell.alive = cell.nextAliveStatus
		}
	}
}

func NewGame(rows int, cols int) *game {
	game := game{
		rows: rows,
		cols: cols,
	}

	// Create cells acording to the given rows and columns
	for r := 0; r <= rows-1; r++ {
		var newRow []*cell
		for c := 0; c <= cols-1; c++ {
			newRow = append(newRow, &cell{x: c, y: r, alive: false})
		}
		game.cellsMatrix = append(game.cellsMatrix, newRow)
	}

	game.genRandomState()

	return &game
}


func (cell *cell) setIsAive(newState bool) {
	cell.alive = newState
}

func renderGame(game *game) {
	// Prints the char if the cell is alive
	for row, _ := range game.cellsMatrix {
		for _, cell := range game.cellsMatrix[row] {
			if cell.alive {
				fmt.Printf("#")  
			} else {
				fmt.Printf(" ")  
			}
		}
		fmt.Printf("\n")  
	}
}

func main() {
	rows := 30
	cols := 50

	game := NewGame(rows, cols)

	for {
		fmt.Print("\033[H\033[2J")
		renderGame(game)
		game.nextGeneration()
		time.Sleep(1000 * time.Millisecond)
	} 
}
