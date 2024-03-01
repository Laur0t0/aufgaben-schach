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

	// Bedingungen:
	// 	position(Zeile) > 0
	// 	position(Zeile) < Länge der Liste(horizontal)
	// 	position(Reihe) > 0
	// 	position(Reihe) < Länge der Liste (vertikal)
	// 	Position im Bord ist lehr

	return pos.Row >= 0 && pos.Row < len(b) && pos.Col >= 0 && pos.Col < len(b[0]) && b[pos.Row][pos.Col] == " "

	//oder mit if

	// if pos.Row >= 0 && pos.Row < len(b) && pos.Col >= 0 && pos.Col < len(b[0]) && b[pos.Row][pos.Col] == " " {
	// 	return true
	// } else {
	// 	return false
	// }

	//oder invertiert mit if

	// if pos.Row < 0 || pos.Row >= len(b) || pos.Col < 0 || pos.Col >= len(b[0]) || b[pos.Row][pos.Col] != " " {
	// 	return false
	// }

	// return true
}

// KnightNeighbours erwartet eine Spielfeldposition und liefert eine Liste mit
// allen von dort mittels einer Springer-Bewegung erreichbaren Positionen,
func KnightNeighbours(pos BoardPos) []BoardPos {
	//hartcodierte Liste, Positionen relativ zu pos.Row und pos.Col gespeichert
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
