package queens

import "chesspuzzles/board"

// QueenAllowed erwartet ein Spielfeld und eine Position.
// Liefert true, falls es erlaubt ist, dort eine Dame zu setzen.
func QueenAllowed(b board.Board, row, col int) bool {
	/* Hinweis:
	 * Verwenden Sie die Funktionen RowEmpty, ColumnEmpty, ... aus dem Package board.
	 */
	return board.RowEmpty(b, row) &&
		board.ColumnEmpty(b, col) &&
		board.DiagDownRightEmpty(b, col-row) &&
		board.DiagUpRightEmpty(b, col-(len(b)-1-row))
}
