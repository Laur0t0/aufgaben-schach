package board

// OnlySpaces erwartet eine Liste von Strings und liefert true,
// wenn alle Strings Leerzeichen sind, sonst false.
func OnlySpaces(list []string) bool {
	for _, v := range list {
		if v != " " {
			return false
		}
	}
	return true
}

// ContainsQueen erwarte eine Liste von Strings und liefert true,
// wenn darin mindestens ein "Q" enthalten ist, sonst false.
func ContainsQueen(list []string) bool {
	for _, v := range list {
		if v == "Q" {
			return true
		}
	}
	return false
}

// RowEmpty erwartet ein Spielfeld und eine Zeilennummer row.
// Liefert true, wenn die row-te Zeile leer ist, sonst false.
// Liefert true, wenn die Zeile nicht existiert.
func RowEmpty(board Board, row int) bool {
	return OnlySpaces(GetRow(board, row))
}

// ColumnEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, wenn die col-te Spalte leer ist, sonst false.
// Liefert true, wenn die Spalte nicht existiert.
func ColumnEmpty(board Board, col int) bool {
	return OnlySpaces(GetColumn(board, col))
}

// DiagDownRightEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und Zeile 0 beginnt,
// nur Leerzeichen enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func DiagDownRightEmpty(board Board, col int) bool {
	return OnlySpaces(GetDiagDownRight(board, col))
}

// DiagUpRightEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und der letzten Zeile beginnt,
// nur Leerzeichen enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func DiagUpRightEmpty(board Board, col int) bool {
	return OnlySpaces(GetDiagUpRight(board, col))
}

// RowContainsQueen erwartet ein Spielfeld und eine Zeilennummer row.
// Liefert true, falls die row-te Zeile mindestens ein "Q" enthält,
// sonst false.
// Liefert false, falls die Zeile nicht existiert.
func RowContainsQueen(board Board, row int) bool {
	return ContainsQueen(GetRow(board, row))
}

// ColumnContainsQueen erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die col-te Spalte mindestens ein "Q" enthält,
// sonst false.
// Liefert false, falls die Spalte nicht existiert.
func ColumnContainsQueen(board Board, col int) bool {
	return ContainsQueen(GetColumn(board, col))
}

// DiagDownRightContainsQueen erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und Zeile 0 beginnt,
// mindestens ein "Q" enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func DiagDownRightContainsQueen(board Board, col int) bool {
	return ContainsQueen(GetDiagDownRight(board, col))
}

// DiagUpRightContainsQueen erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und der letzten Zeile beginnt,
// mindestens ein "Q" enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func DiagUpRightContainsQueen(board Board, col int) bool {
	return ContainsQueen(GetDiagUpRight(board, col))
}
