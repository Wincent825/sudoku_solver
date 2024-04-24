package main

import (
	"fmt"
	"os"
	"time"
)

type Cell struct {
	row, column int
	value       int
	empty       bool
}

type Board struct {
	cells [][]Cell
}

func main() {
	startTime := time.Now()
	board := parseArgs()

	//board.printBoard()

	var endTime time.Duration
	if board.solveSudoku(board) {
		board.printBoard()
		if controlBoard(board) {
			endTime = time.Since(startTime)
		} else {
			fmt.Println("solution is invalid, please check the board and try again")
			endTime = time.Since(startTime)
		}
	} else {
		fmt.Println("No solution found")
		endTime = time.Since(startTime)
	}
	fmt.Println("Execution time: ", endTime)
}

func (board *Board) printBoard() {
	// Print the sudoku board cell by cell
	BoardCells := board.cells

	for i, row := range BoardCells {
		// Print a newline after every 3 rows
		if i == 3 || i == 6 {
			fmt.Println("------+-------+------")
		}

		for j, cell := range row {
			if cell.empty {
				fmt.Print(".")
			} else {
				fmt.Print(cell.value)
			}

			if j == 8 {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}

			// Print a newline after every 3 rows
			if j == 2 || j == 5 {
				fmt.Print("|")
			}
		}
	}
}

func parseArgs() Board {
	args := os.Args[1:]

	if len(args) != 9 {
		fmt.Println("Error: Invalid number of arguments")
		os.Exit(1)
	}

	for _, arg := range args {
		if len(arg) != 9 {
			fmt.Println("Error: Invalid argument length")
			os.Exit(1)
		}

		for _, char := range arg {
			if char != '.' && (char < '1' || char > '9') {
				fmt.Println("Error: Invalid character in argument")
				os.Exit(1)
			}
		}
	}

	// If the arguments are valid, return a 9x9 matrix
	// representing the sudoku board
	// Each board cell is represented by a number from 0-9
	// boolean values are used to represent empty cells
	// make sure there are no duplicate numbers in the same row, column or 3x3 subgrid
	cells := make([][]Cell, 9)

	for i, arg := range args {
		cells[i] = make([]Cell, 9)

		for j, char := range arg {
			if char == '.' {
				cells[i][j] = Cell{row: i, column: j, value: 0, empty: true}
			} else {
				cells[i][j] = Cell{row: i, column: j, value: int(char - '0'), empty: false}
			}
		}
	}

	if len(cells) != 9 {
		fmt.Println("Error: Invalid number of rows")
		os.Exit(1)
	}

	for _, row := range cells {
		if len(row) != 9 {
			fmt.Println("Error: Invalid number of columns")
			os.Exit(1)
		}
	}

	gameboard := Board{cells: cells}

	return gameboard
}

func (b *Board) solveSudoku(board Board) bool {
	// find the first empty cell
	// try to fill it with a number from 1-9
	// if the number is valid, move to the next cell

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board.cells[row][col].empty {
				for num := 1; num <= 9; num++ {
					//board.printBoard()
					if isValidMove(board, row, col, num) {
						board.cells[row][col].value = num
						board.cells[row][col].empty = false

						// recursively solve the remaining board
						if board.solveSudoku(board) {
							return true
						}

						// if the current assignment does not lead to a solution, backtrack
						board.cells[row][col].value = 0
						board.cells[row][col].empty = true
					}
				}
				// if no valid number can be assigned to the current cell, return false
				return false
			}
		}
	}

	// if no empty cell is found, the puzzle is solved
	return true
}

func isValidMove(board Board, row, column, num int) bool {
	// check if the number is valid in the given row and column
	for i := 0; i < 9; i++ {
		if board.cells[row][i].value == num || board.cells[i][column].value == num {
			return false
		}
	}

	// check if the number is valid in the 3x3 subgrid
	startRow := row - row%3
	startColumn := column - column%3

	for i := startRow; i < startRow+3; i++ {
		for j := startColumn; j < startColumn+3; j++ {
			if board.cells[i][j].value == num {
				return false
			}
		}
	}

	return true
}

func controlBoard(board Board) bool {
	// check if the board is valid and has no duplicates
	// in the same row, column or 3x3 subgrid

	// check rows, columns and 3x3 subgrids
	if !checkRow(board) || !checkColumn(board) || !checkSubgrid(board) {
		return false
	}

	return true
}

func checkRow(board Board) bool {
	// check rows
	for i := 0; i < 9; i++ {
		row := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if board.cells[i][j].value != 0 {
				if row[board.cells[i][j].value] {
					return false
				}
				row[board.cells[i][j].value] = true
			}
		}
	}
	return true
}

func checkColumn(board Board) bool {
	// check columns
	for i := 0; i < 9; i++ {
		column := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if board.cells[j][i].value != 0 {
				if column[board.cells[j][i].value] {
					return false
				}
				column[board.cells[j][i].value] = true
			}
		}
	}
	return true
}

func checkSubgrid(board Board) bool {
	// check 3x3 subgrids
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			subgrid := make(map[int]bool)
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					if board.cells[k][l].value != 0 {
						if subgrid[board.cells[k][l].value] {
							return false
						}
						subgrid[board.cells[k][l].value] = true
					}
				}
			}
		}
	}
	return true
}
