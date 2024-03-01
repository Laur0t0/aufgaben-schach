package main

import (
	"chesspuzzles/board"
	"chesspuzzles/knights"
	"fmt"
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
	// 1. Wenn n = negativ oder zu groß: Spiel per Definitionem gelöst.
	// 2. Wenn Position ungültig oder bereits vergeben (verbotener Zug): Spiel nicht lösbar.

	if n <= 0 || len(b) == 0 || len(b[0]) == 0 || n > len(b)*len(b[0]) {
		return true
	}

	if !knights.KnightAllowed(b, pos) {
		return false
	}

	//n muss auf das aktuelle Feld geschrieben werden1
	b[pos.Row][pos.Col] = fmt.Sprintf("%d", n)

	//Prüfen für jede Nachbarposition (mit knightNeighbours), ob Spiel ab dort mit n+1 zu lösen geht
	for _, p := range knights.KnightNeighbours(pos) {
		if SolveKnights(b, p, n+1) {
			return true
		}
	}

	//wenn das Spiel durch das nächste Feld nicht gelöst werden kann, dann muss dieser gelöscht werden b[pos.Row][pos.Col] = fmt.Sprintf("%d", n)

	b[pos.Row][pos.Col] = " "

	return false
}
