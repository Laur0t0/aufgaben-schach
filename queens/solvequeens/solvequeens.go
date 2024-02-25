package main

import (
	"chesspuzzles/board"
	"chesspuzzles/queens"
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
	/* Hinweis:
	 * Rekursionsanker: Für ungültige Zeilennummern ist das Spiel per Definitionem gelöst.
	 */
	if row < 0 || row >= len(b) {
		return true
	}

	/* Hinweis:
	* Laufen Sie mit einer Schleife durch die aktuelle Zeile.
	* Falls ein Zug an der jeweiligen Position erlaubt ist, setzen Sie dort eine Dame
	* und machen Sie rekursiv in der nächsten Zeile weiter.
	 */
	for col := range b[row] {
		if queens.QueenAllowed(b, row, col) {
			b[row][col] = "Q"
			if SolveQueens(b, row+1) {
				return true
			}
			b[row][col] = " "
		}
	}

	/* Hinweis:
	* Ist das Spiel bis hier (nach der Schleife) nicht gelöst,
	* ist es mit dem gegebenen Board gar nicht lösbar.
	 */
	return false
}
