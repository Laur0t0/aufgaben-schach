package board

import "fmt"

func ExampleOnlySpaces() {
	fmt.Println(OnlySpaces([]string{" ", " ", " "}))
	fmt.Println(OnlySpaces([]string{" ", "X", " "}))
	fmt.Println(OnlySpaces([]string{}))

	// Output:
	// true
	// false
	// true
}

func ExampleContainsQueen() {
	fmt.Println(ContainsQueen([]string{" ", " ", " "}))
	fmt.Println(ContainsQueen([]string{" ", "Q", " "}))
	fmt.Println(ContainsQueen([]string{}))

	// Output:
	// false
	// true
	// false
}

func ExampleRowEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", "Q", " "}
	board[2] = []string{" ", " ", " "}

	fmt.Println(RowEmpty(board, 0))
	fmt.Println(RowEmpty(board, 1))
	fmt.Println(RowEmpty(board, 2))
	fmt.Println(RowEmpty(board, -1))
	fmt.Println(RowEmpty(board, 3))

	// Output:
	// true
	// false
	// true
	// true
	// true
}

func ExampleColumnEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{" ", " ", " "}

	fmt.Println(ColumnEmpty(board, 0))
	fmt.Println(ColumnEmpty(board, 1))
	fmt.Println(ColumnEmpty(board, 2))
	fmt.Println(ColumnEmpty(board, -1))
	fmt.Println(ColumnEmpty(board, 3))

	// Output:
	// true
	// true
	// false
	// true
	// true
}

func ExampleDiagDownRightEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{"Q", "Q", " "}

	fmt.Println(DiagDownRightEmpty(board, -3))
	fmt.Println(DiagDownRightEmpty(board, -2))
	fmt.Println(DiagDownRightEmpty(board, -1))
	fmt.Println(DiagDownRightEmpty(board, 0))
	fmt.Println(DiagDownRightEmpty(board, 1))
	fmt.Println(DiagDownRightEmpty(board, 2))
	fmt.Println(DiagDownRightEmpty(board, 3))

	// Output:
	// true
	// false
	// false
	// true
	// false
	// true
	// true
}

func ExampleDiagUpRightEmpty() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{"Q", "Q", " "}

	fmt.Println(DiagUpRightEmpty(board, -3))
	fmt.Println(DiagUpRightEmpty(board, -2))
	fmt.Println(DiagUpRightEmpty(board, -1))
	fmt.Println(DiagUpRightEmpty(board, 0))
	fmt.Println(DiagUpRightEmpty(board, 1))
	fmt.Println(DiagUpRightEmpty(board, 2))
	fmt.Println(DiagUpRightEmpty(board, 3))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// true
	// true
}

func ExampleRowContainsQueen() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", "Q", " "}
	board[2] = []string{" ", " ", " "}

	fmt.Println(RowContainsQueen(board, 0))
	fmt.Println(RowContainsQueen(board, 1))
	fmt.Println(RowContainsQueen(board, 2))
	fmt.Println(RowContainsQueen(board, -1))
	fmt.Println(RowContainsQueen(board, 3))

	// Output:
	// false
	// true
	// false
	// false
	// false
}

func ExampleColumnContainsQueen() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{" ", " ", " "}

	fmt.Println(ColumnContainsQueen(board, 0))
	fmt.Println(ColumnContainsQueen(board, 1))
	fmt.Println(ColumnContainsQueen(board, 2))
	fmt.Println(ColumnContainsQueen(board, -1))
	fmt.Println(ColumnContainsQueen(board, 3))

	// Output:
	// false
	// false
	// true
	// false
	// false
}

func ExampleDiagDownRightContainsQueen() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{"Q", "Q", " "}

	fmt.Println(DiagDownRightContainsQueen(board, -3))
	fmt.Println(DiagDownRightContainsQueen(board, -2))
	fmt.Println(DiagDownRightContainsQueen(board, -1))
	fmt.Println(DiagDownRightContainsQueen(board, 0))
	fmt.Println(DiagDownRightContainsQueen(board, 1))
	fmt.Println(DiagDownRightContainsQueen(board, 2))
	fmt.Println(DiagDownRightContainsQueen(board, 3))

	// Output:
	// false
	// true
	// true
	// false
	// true
	// false
	// false
}

func ExampleDiagUpRightContainsQueen() {
	board := EmptyBoard(3)
	board[0] = []string{" ", " ", " "}
	board[1] = []string{" ", " ", "Q"}
	board[2] = []string{"Q", "Q", " "}

	fmt.Println(DiagUpRightContainsQueen(board, -3))
	fmt.Println(DiagUpRightContainsQueen(board, -2))
	fmt.Println(DiagUpRightContainsQueen(board, -1))
	fmt.Println(DiagUpRightContainsQueen(board, 0))
	fmt.Println(DiagUpRightContainsQueen(board, 1))
	fmt.Println(DiagUpRightContainsQueen(board, 2))
	fmt.Println(DiagUpRightContainsQueen(board, 3))

	// Output:
	// false
	// false
	// false
	// true
	// true
	// false
	// false
}
