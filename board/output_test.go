package board

import "fmt"

func ExampleEmptyBoard() {
	fmt.Println(EmptyBoard(3))

	// Output:
	// [[     ] [     ] [     ]]
}

func ExamplePrintBoard() {
	board := EmptyBoard(3)
	board[1][1] = "Q"
	PrintBoard(board)

	// Output:
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   | Q |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
}

func ExamplePrintBoard_wide() {
	board := EmptyBoard(3)
	board[1][1] = "QQ"
	board[0][1] = "Q"
	PrintBoard(board)

	// Output:
	// +----+----+----+
	// |    |  Q |    |
	// +----+----+----+
	// |    | QQ |    |
	// +----+----+----+
	// |    |    |    |
	// +----+----+----+
}
