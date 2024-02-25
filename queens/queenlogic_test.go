package queens

import (
	"fmt"

	"chesspuzzles/board"
)

func ExampleQueenAllowed() {
	b := board.EmptyBoard(5)
	b[0] = []string{" ", "Q", " ", " ", " "}
	b[1] = []string{" ", " ", " ", " ", " "}
	b[2] = []string{" ", " ", " ", " ", " "}
	b[3] = []string{" ", " ", " ", " ", " "}
	b[4] = []string{" ", " ", "Q", " ", " "}

	// Erlaubte Positionen:
	fmt.Println(QueenAllowed(b, 1, 3))
	fmt.Println(QueenAllowed(b, 1, 4))
	fmt.Println(QueenAllowed(b, 3, 0))
	fmt.Println()

	// Verbotene Positionen:
	// Spalte 0
	fmt.Println(QueenAllowed(b, 0, 0))
	fmt.Println(QueenAllowed(b, 0, 1))
	fmt.Println(QueenAllowed(b, 0, 2))
	fmt.Println(QueenAllowed(b, 0, 3))
	fmt.Println(QueenAllowed(b, 0, 4))
	fmt.Println()

	// Spalte 1
	fmt.Println(QueenAllowed(b, 1, 0))
	fmt.Println(QueenAllowed(b, 1, 1))
	fmt.Println(QueenAllowed(b, 1, 2))
	fmt.Println()

	// Spalte 2
	fmt.Println(QueenAllowed(b, 2, 0))
	fmt.Println(QueenAllowed(b, 2, 1))
	fmt.Println(QueenAllowed(b, 2, 2))
	fmt.Println(QueenAllowed(b, 2, 3))
	fmt.Println(QueenAllowed(b, 2, 4))
	fmt.Println()

	// Spalte 3
	fmt.Println(QueenAllowed(b, 3, 1))
	fmt.Println(QueenAllowed(b, 3, 2))
	fmt.Println(QueenAllowed(b, 3, 3))
	fmt.Println(QueenAllowed(b, 3, 4))
	fmt.Println()

	// Spalte 4
	fmt.Println(QueenAllowed(b, 4, 0))
	fmt.Println(QueenAllowed(b, 4, 1))
	fmt.Println(QueenAllowed(b, 4, 2))
	fmt.Println(QueenAllowed(b, 4, 3))
	fmt.Println(QueenAllowed(b, 4, 4))

	// Output:
	// true
	// true
	// true
	//
	// false
	// false
	// false
	// false
	// false
	//
	// false
	// false
	// false
	//
	// false
	// false
	// false
	// false
	// false
	//
	// false
	// false
	// false
	// false
	//
	// false
	// false
	// false
	// false
	// false
}
