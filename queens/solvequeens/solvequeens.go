package main

import (
	"chesspuzzles/board"
)

func main() {
	for n := 1; n <= 10; n++ {
		b := board.EmptyBoard(n)
		SolveQueens(b, 0)
		board.PrintBoard(b)
	}
}

// SolveQueens erwartet ein Spielfeld und eine Zeilennummer row.
// Versucht, das Spiel ab Zeile row zu lösen.
// Liefert true, falls dies gelingt und setzt die entsprechenden Damen.
func SolveQueens(b board.Board, row int) bool {
	// TODO
	return false
}
