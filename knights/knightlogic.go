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
	// TODO
	return false
}

// KnightNeighbours erwartet eine Spielfeldposition und liefert eine Liste mit
// allen von dort mittels einer Springer-Bewegung erreichbaren Positionen,
func KnightNeighbours(pos BoardPos) []BoardPos {
	// TODO
	return []BoardPos{}
}
