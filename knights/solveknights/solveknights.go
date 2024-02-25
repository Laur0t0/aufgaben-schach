package main

import (
	"chesspuzzles/board"
	"chesspuzzles/knights"
)

func main() {
	b := board.EmptyBoard(8)
	SolveKnights(b, knights.BoardPos{Row: 0, Col: 0}, 1)
	board.PrintBoard(b)
}

// SolveKnights erwartet ein Spielfeld und eine Position, sowie eine laufende Nummer n.
// Versucht, das Spiel ab der Nummer n zu lösen, indem es n an die aktuelle Position
// schreibt und dann rekursiv mit allen für den Springer erreichbaren Feldern weitermacht.
func SolveKnights(b board.Board, pos knights.BoardPos, n int) bool {
	// TODO

	return false
}
