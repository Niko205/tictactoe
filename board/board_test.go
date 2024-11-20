package board

import "fmt"

func ExampleNew() {
	board := New(3, 3, "*")

	for _, row := range board {
		fmt.Println(row)
	}

	// Output:
	// | * | * | * |
	// | * | * | * |
	// | * | * | * |
}

func ExampleBoard_String() {
	board := New(3, 3, "*")

	fmt.Println(board)

	// Output:
	// +---+---+---+
	// | * | * | * |
	// +---+---+---+
	// | * | * | * |
	// +---+---+---+
	// | * | * | * |
	// +---+---+---+
}

func ExampleBoard_Row() {
	board := New(3, 4, "*")

	fmt.Println(board.Row(1))

	// Output:
	// | * | * | * | * |
}

func ExampleBoard_Col() {
	board := New(3, 4, "*")

	fmt.Println(board.Col(1))

	// Output:
	// | * | * | * |
}

func ExampleBoard_Set() {
	board := New(3, 3, "*")

	board.Set(1, 1, "X")

	fmt.Println(board)

	// Output:
	// +---+---+---+
	// | * | * | * |
	// +---+---+---+
	// | * | X | * |
	// +---+---+---+
	// | * | * | * |
	// +---+---+---+
}

func ExampleBoard_Diag() {
	board := Board{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	fmt.Println(board.Diag(0))
	fmt.Println(board.Diag(1))

	// Output:
	// | 1 | 5 | 9 |
	// | 3 | 5 | 7 |
}

func ExampleBoard_Full() {
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

	fmt.Println(board1.Full())
	fmt.Println(board2.Full())

	// Output:
	// true
	// false
}

func ExampleBoard_RowContainsOnly() {
	board := Board{
		{"X", "X", " "},
		{"O", "O", "O"},
		{" ", " ", " "},
	}

	fmt.Println(board.RowContainsOnly(0, "X"))
	fmt.Println(board.RowContainsOnly(0, "O"))
	fmt.Println(board.RowContainsOnly(0, " "))

	fmt.Println(board.RowContainsOnly(1, "X"))
	fmt.Println(board.RowContainsOnly(1, "O"))
	fmt.Println(board.RowContainsOnly(1, " "))

	fmt.Println(board.RowContainsOnly(2, "X"))
	fmt.Println(board.RowContainsOnly(2, "O"))
	fmt.Println(board.RowContainsOnly(2, " "))

	// Output:
	// false
	// false
	// false
	// false
	// true
	// false
	// false
	// false
	// true
}

func ExampleBoard_ColContainsOnly() {
	board := Board{
		{"X", "X", "X"},
		{"O", "O", "X"},
		{"X", "O", "X"},
	}

	fmt.Println(board.ColContainsOnly(0, "X"))
	fmt.Println(board.ColContainsOnly(0, "O"))
	fmt.Println(board.ColContainsOnly(0, " "))

	fmt.Println(board.ColContainsOnly(1, "X"))
	fmt.Println(board.ColContainsOnly(1, "O"))
	fmt.Println(board.ColContainsOnly(1, " "))

	fmt.Println(board.ColContainsOnly(2, "X"))
	fmt.Println(board.ColContainsOnly(2, "O"))
	fmt.Println(board.ColContainsOnly(2, " "))

	// Output:
	// false
	// false
	// false
	// false
	// false
	// false
	// true
	// false
	// false
}

func ExampleBoard_DiagContainsOnly() {
	board := Board{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}

	fmt.Println(board.DiagContainsOnly(0, "X"))
	fmt.Println(board.DiagContainsOnly(0, "O"))
	fmt.Println(board.DiagContainsOnly(0, " "))

	fmt.Println(board.DiagContainsOnly(1, "X"))
	fmt.Println(board.DiagContainsOnly(1, "O"))
	fmt.Println(board.DiagContainsOnly(1, " "))

	// Output:
	// true
	// false
	// false
	// true
	// false
	// false
}
