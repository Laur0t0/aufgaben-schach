package board

type Board [][]string

// EmptyBoard liefert ein neues, leeres Spielfeld der Größe nxn.
func EmptyBoard(n int) Board {
	board := Board{}

	for i := 0; i < n; i++ {
		line := []string{}
		for i := 0; i < n; i++ {
			line = append(line, " ")
		}
		board = append(board, line)
	}

	return board
}

// GetRow erwartet ein Spielfeld und eine Zeilennummer row.
// Liefert die row-te Zeile des Spielfelds.
// Liefert eine leere Liste, falls die Zeile nicht existiert.
func GetRow(board Board, row int) []string {
	if row < 0 || row >= len(board) {
		return []string{}
	}
	return board[row]
}

// GetColumn erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert die col-te Spalte des Spielfelds.
// Liefert eine leere Liste, falls die Spalte nicht existiert.
func GetColumn(board Board, col int) []string {
	if col < 0 || col >= len(board[0]) {
		return []string{}
	}
	result := []string{}

	for _, row := range board {
		result = append(result, row[col])
	}
	return result
}

// GetDiagDownRight erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert die Diagonale, die in Spalte col und Zeile 0 beginnt
// und die von dort aus nach rechts unten verläuft.
// Für ungültige Spaltennummern wird ggf. eine Teil-Diagonale geliefert.
func GetDiagDownRight(board Board, col int) []string {
	result := []string{}

	for i := 0; i < len(board); i++ {
		r, c := i, col+i
		if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) {
			result = append(result, board[r][c])
		}
	}
	return result
}

// GetDiagUpRight erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert die Diagonale, die in Spalte col und der letzten Zeile beginnt
// und die von dort aus nach rechts oben verläuft.
// Für ungültige Spaltennummern wird ggf. eine Teil-Diagonale geliefert.
func GetDiagUpRight(board Board, col int) []string {
	result := []string{}
	for i := 0; i < len(board); i++ {
		r, c := len(board)-1-i, col+i
		if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) {
			result = append(result, board[r][c])
		}
	}
	return result
}
