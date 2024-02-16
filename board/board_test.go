package board

import "fmt"

func ExampleGetRow() {
	board := EmptyBoard(3)
	board[0] = []string{"A", "B", "C"}
	board[1] = []string{"D", "E", "F"}
	board[2] = []string{"G", "H", "I"}

	fmt.Println(GetRow(board, 0))
	fmt.Println(GetRow(board, 1))
	fmt.Println(GetRow(board, 2))

	fmt.Println(GetRow(board, -1))
	fmt.Println(GetRow(board, 3))

	// Output:
	// [A B C]
	// [D E F]
	// [G H I]
	// []
	// []
}

func ExampleGetColumn() {
	board := EmptyBoard(3)
	board[0] = []string{"A", "B", "C"}
	board[1] = []string{"D", "E", "F"}
	board[2] = []string{"G", "H", "I"}

	fmt.Println(GetColumn(board, 0))
	fmt.Println(GetColumn(board, 1))
	fmt.Println(GetColumn(board, 2))

	fmt.Println(GetColumn(board, -1))
	fmt.Println(GetColumn(board, 3))

	// Output:
	// [A D G]
	// [B E H]
	// [C F I]
	// []
	// []
}

func ExampleGetDiagDownRight() {
	board := EmptyBoard(3)
	board[0] = []string{"A", "B", "C"}
	board[1] = []string{"D", "E", "F"}
	board[2] = []string{"G", "H", "I"}

	fmt.Println(GetDiagDownRight(board, -3))
	fmt.Println(GetDiagDownRight(board, -2))
	fmt.Println(GetDiagDownRight(board, -1))
	fmt.Println(GetDiagDownRight(board, 0))
	fmt.Println(GetDiagDownRight(board, 1))
	fmt.Println(GetDiagDownRight(board, 2))
	fmt.Println(GetDiagDownRight(board, 3))

	// Output:
	// []
	// [G]
	// [D H]
	// [A E I]
	// [B F]
	// [C]
	// []
}

func ExampleGetDiagUpRight() {
	board := EmptyBoard(3)
	board[0] = []string{"A", "B", "C"}
	board[1] = []string{"D", "E", "F"}
	board[2] = []string{"G", "H", "I"}

	fmt.Println(GetDiagUpRight(board, -3))
	fmt.Println(GetDiagUpRight(board, -2))
	fmt.Println(GetDiagUpRight(board, -1))
	fmt.Println(GetDiagUpRight(board, 0))
	fmt.Println(GetDiagUpRight(board, 1))
	fmt.Println(GetDiagUpRight(board, 2))
	fmt.Println(GetDiagUpRight(board, 3))

	// Output:
	// []
	// [A]
	// [D B]
	// [G E C]
	// [H F]
	// [I]
	// []
}
