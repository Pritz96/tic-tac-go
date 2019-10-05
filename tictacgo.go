package main

import "fmt"

var grid = [3][3]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
var row int
var col int
var playerStamp string

func main() {
	i := 0
	fmt.Println(printGrid(grid))

turns:
	for {
		if i%2 == 0 {
			playerStamp = "X"
		} else {
			playerStamp = "O"
		}

	validPos:
		for {

		validRow:
			for {
				fmt.Println("Player " + playerStamp + " enter your row: ")
				fmt.Scanln(&row)
				if row < 0 || row >= 3 {
					fmt.Println("Row is not 0, 1 or 2, please re-enter row")
				} else {
					break validRow
				}
			}

		validCol:
			for {
				fmt.Println("Player " + playerStamp + " enter your col: ")
				fmt.Scanln(&col)
				if col < 0 || col >= 3 {
					fmt.Println("Col is not 0, 1 or 2, please re-enter col")
				} else {
					break validCol
				}
			}

			if isPositionAlreadyFilled(grid, row, col) {
				fmt.Println("Position is already filled in the grid, re-enter row and col")
			} else {
				grid[row][col] = playerStamp
				fmt.Println(printGrid(grid))
				break validPos
			}

		}
		if hasAnyoneWonYet(grid) {
			fmt.Println("Player " + playerStamp + " has won")
			break turns
		} else if allPositionsFilled(grid) {
			fmt.Println("The game is tied")
			break turns
		} else {
			i = i + 1
		}
	}

}

func printGrid(grid [3][3]string) string {
	visualGrid := fmt.Sprintf("|-%v-|-%v-|-%v-|\n|-%v-|-%v-|-%v-|\n|-%v-|-%v-|-%v-|", grid[0][0], grid[0][1], grid[0][2], grid[1][0], grid[1][1], grid[1][2], grid[2][0], grid[2][1], grid[2][2])
	return visualGrid
}

func isPositionAlreadyFilled(grid [3][3]string, row, col int) bool {
	if grid[row][col] == "X" || grid[row][col] == "O" {
		return true
	}
	return false
}

func allPositionsFilled(grid [3][3]string) bool {
	count := 0
	for j := 0; j < 3; j++ {
		if isPositionAlreadyFilled(grid, 0, j) && isPositionAlreadyFilled(grid, 1, j) && isPositionAlreadyFilled(grid, 2, j) {
			count = count + 1
		}
	}
	if count == 3 {
		return true
	}
	return false
}

func hasAnyoneWonYet(grid [3][3]string) bool {
	if checkAllRows(grid) || checkAllColumns(grid) || checkDiagonals(grid) {
		return true
	}
	return false
}

func checkAllRows(grid [3][3]string) bool {
	row1 := grid[0][0] == grid[0][1] && grid[0][1] == grid[0][2]
	row2 := grid[1][0] == grid[1][1] && grid[1][1] == grid[1][2]
	row3 := grid[2][0] == grid[2][1] && grid[2][1] == grid[2][2]

	if row1 && sameShape(grid[0][0], grid[0][1], grid[0][2]) || row2 && sameShape(grid[1][0], grid[1][1], grid[1][2]) || row3 && sameShape(grid[2][0], grid[2][1], grid[2][2]) {
		return true
	}
	return false
}

func checkAllColumns(grid [3][3]string) bool {
	col1 := grid[0][0] == grid[1][0] && grid[1][0] == grid[2][0]
	col2 := grid[0][1] == grid[1][1] && grid[1][1] == grid[2][1]
	col3 := grid[0][2] == grid[1][2] && grid[1][2] == grid[2][2]

	if col1 && sameShape(grid[0][0], grid[1][0], grid[2][0]) || col2 && sameShape(grid[0][1], grid[1][1], grid[2][1]) || col3 && sameShape(grid[0][2], grid[1][2], grid[2][2]) {
		return true
	}
	return false
}

func checkDiagonals(grid [3][3]string) bool {
	diag1 := grid[0][0] == grid[1][1] && grid[1][1] == grid[2][2]
	diag2 := grid[0][2] == grid[1][1] && grid[1][1] == grid[2][0]

	if diag1 && sameShape(grid[0][0], grid[1][1], grid[2][2]) || diag2 && sameShape(grid[0][2], grid[1][1], grid[2][0]) {
		return true
	}
	return false
}

func sameShape(a, b, c string) bool {
	if a == b && b == c && c == "X" || a == b && b == c && c == "O" {
		return true
	}
	return false
}
