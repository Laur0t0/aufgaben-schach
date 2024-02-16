package board

import (
	"fmt"
	"strings"
)

// LargestWidth ermittelt die maximale Breite eines
// Eintrags auf dem Spielfeld.
func LargestWidth(board Board) int {
	result := 0
	for _, row := range board {
		for _, col := range row {
			if len(col) > result {
				result = len(col)
			}
		}
	}
	return result
}

// PrintBoard erwartet ein Spielfeld und gibt es so auf die Konsole aus,
// dass es für Menschen erkennbar ist.
func PrintBoard(board Board) {
	n := len(board[0])
	lw := LargestWidth(board)

	divpart := fmt.Sprintf("+-%s-", strings.Repeat("-", lw))
	divider := strings.Repeat(divpart, n) + "+"

	for i := 0; i < n; i++ {
		fmt.Println(divider)
		row := board[i]
		for j := 0; j < n; j++ {
			padding := strings.Repeat(" ", lw-len(row[j]))
			fmt.Printf("| %s%s ", padding, row[j])
		}
		fmt.Println("|")
	}
	fmt.Println(divider)
}
