package knights

import "chesspuzzles/board"

// BoardPos repräsentiert eine Position auf dem Spielfeld.
type BoardPos struct {
	Row int
	Col int
}

// KnightAllowed erwartet ein Spielfeld und eine Position.
// Liefert true, falls auf dieses Feld ein Springer gesetzt werden darf.
func KnightAllowed(b board.Board, pos BoardPos) bool {
	/* Hinweis Es gibt hier keine Einschränkungen,
	 * außer dass die Position gültig sein muss
	 * und das Feld nicht besetzt sein darf.
	 */
	return pos.Row >= 0 && pos.Row < len(b) && pos.Col >= 0 && pos.Col < len(b[0]) && b[pos.Row][pos.Col] == " "
}

// KnightNeighbours erwartet eine Spielfeldposition und liefert eine Liste mit
// allen von dort mittels einer Springer-Bewegung erreichbaren Positionen,
func KnightNeighbours(pos BoardPos) []BoardPos {
	/* Hinweis:
	 * Am einfachsten hartcodieren Sie eine Liste, in der die Positionen
	 * relativ zu pos.Row und pos.Col gespeichert sind.
	 */
	return []BoardPos{
		{pos.Row - 1, pos.Col - 2},
		{pos.Row - 1, pos.Col + 2},
		{pos.Row + 1, pos.Col - 2},
		{pos.Row + 1, pos.Col + 2},
		{pos.Row - 2, pos.Col - 1},
		{pos.Row - 2, pos.Col + 1},
		{pos.Row + 2, pos.Col - 1},
		{pos.Row + 2, pos.Col + 1},
	}
}
