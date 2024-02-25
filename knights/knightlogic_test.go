package knights

import (
	"chesspuzzles/board"
	"fmt"
)

func ExampleKnightAllowed() {
	b := board.EmptyBoard(4)
	b[0] = []string{"1", " ", " ", " "}
	b[1] = []string{" ", " ", "2", " "}
	b[2] = []string{"3", " ", " ", " "}
	b[3] = []string{" ", " ", "4", " "}

	fmt.Println(KnightAllowed(b, BoardPos{0, 0}))
	fmt.Println(KnightAllowed(b, BoardPos{0, 1}))
	fmt.Println(KnightAllowed(b, BoardPos{0, 2}))
	fmt.Println(KnightAllowed(b, BoardPos{0, 3}))
	fmt.Println()
	fmt.Println(KnightAllowed(b, BoardPos{1, 0}))
	fmt.Println(KnightAllowed(b, BoardPos{1, 1}))
	fmt.Println(KnightAllowed(b, BoardPos{1, 2}))
	fmt.Println(KnightAllowed(b, BoardPos{1, 3}))
	fmt.Println()
	fmt.Println(KnightAllowed(b, BoardPos{2, 0}))
	fmt.Println(KnightAllowed(b, BoardPos{2, 1}))
	fmt.Println(KnightAllowed(b, BoardPos{2, 2}))
	fmt.Println(KnightAllowed(b, BoardPos{2, 3}))
	fmt.Println()
	fmt.Println(KnightAllowed(b, BoardPos{3, 0}))
	fmt.Println(KnightAllowed(b, BoardPos{3, 1}))
	fmt.Println(KnightAllowed(b, BoardPos{3, 2}))
	fmt.Println(KnightAllowed(b, BoardPos{3, 3}))

	// Output:
	// false
	// true
	// true
	// true
	//
	// true
	// true
	// false
	// true
	//
	// false
	// true
	// true
	// true
	//
	// true
	// true
	// false
	// true
}

func ExampleKnightNeighbours() {
	p1 := BoardPos{0, 0}
	p2 := BoardPos{3, 3}

	fmt.Println(KnightNeighbours(p1))
	fmt.Println(KnightNeighbours(p2))

	// Output:
	// [{-1 -2} {-1 2} {1 -2} {1 2} {-2 -1} {-2 1} {2 -1} {2 1}]
	// [{2 1} {2 5} {4 1} {4 5} {1 2} {1 4} {5 2} {5 4}]
}
