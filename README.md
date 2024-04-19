# sudoku_solver
This is a simple Sudoku solver implemented in Go. It allows you to input a Sudoku puzzle as a command-line argument and it will attempt to solve it. The solver utilizes backtracking algorithm to recursively find a solution.

## Usage
To use the Sudoku solver, simply run the executable with a string of numbers representing the Sudoku puzzle. Separate each row of the puzzle with double quotes and spaces. Use dots (.) to represent empty cells within each row. The string should have exactly 81 characters in total, representing the 9x9 grid.

For example:

```./sudoku_solver.go ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"```

or:

```go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"```

This will attempt to solve the Sudoku puzzle provided as a command-line argument.

## Sudoku Input Format
The Sudoku input format consists of multiple rows, each enclosed in double quotes and separated by spaces. Within each row, use digits 1-9 to represent filled cells, and dots (.) to represent empty cells.


## Output
The solver will output the solved Sudoku puzzle if a solution is found. If the puzzle has no solution, it will print "No solution found".

## Implementation Details
The solver is implemented using a backtracking algorithm. It iterates through each cell in the Sudoku grid, trying possible numbers from 1 to 9. If a number is found to be valid, it recursively continues to solve the puzzle. If no valid number can be found for a cell, it backtracks to the previous cell and tries a different number.

The validity of a number placement is checked by ensuring there are no duplicates in the same row, column, or 3x3 subgrid.

If the grid has been filled out, a final validator will make sure no duplicates are present.

# Credits
This Sudoku solver was created by Wincent Westerback as a personal project. Feel free to modify and improve it as needed.