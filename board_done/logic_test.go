package board

import "fmt"

func ExampleBoard_PlayerWins() {
	board := Board{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}

	fmt.Println(board.PlayerWins("X"))
	fmt.Println(board.PlayerWins("O"))

	// Output:
	// true
	// false
}

func ExampleBoard_GameOver() {
	board1 := Board{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"O", "X", "O"},
	}

	board2 := Board{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"O", "X", " "},
	}

	board3 := Board{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"O", " ", "X"},
	}

	fmt.Println(board1.GameOver())
	fmt.Println(board2.GameOver())
	fmt.Println(board3.GameOver())

	// Output:
	// true
	// false
	// true
}
