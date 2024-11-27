package ablauf

import (
	"fmt"
	board "tictactoe/board_done"
)

func Ablauf() {
	s := ""
	for {
		b := board.New(3, 3, " ")
		fmt.Println(board.Board.String(b))
		for {
			spielerX(b)
			fmt.Println(board.Board.String(b))
			if b.GameOver() {
				break
			}
			spielerO(b)
			fmt.Println(board.Board.String(b))
			if b.GameOver() {
				break
			}
		}
		fmt.Println("Neues Spiel? J/N")
		fmt.Scan(&s)
		if s == "N" || s == "n" {
			break
		}
	}
}

func spielerX(b board.Board) {
	var x, y int
	for {
		fmt.Println("X: Zeile")
		fmt.Scan(&x)
		fmt.Println("X: Spalte")
		fmt.Scan(&y)
		if b.Empty(x, y) {
			b.Set(x, y, "X")
			break
		} else {
			fmt.Println("Feld ist nicht leer.")
		}
	}
}

func spielerO(b board.Board) {
	var x, y int
	for {
		fmt.Println("O: Zeile")
		fmt.Scan(&x)
		fmt.Println("O: Spalte")
		fmt.Scan(&y)
		if b.Empty(x, y) {
			b.Set(x, y, "O")
			break
		} else {
			fmt.Println("Feld ist nicht leer.")
		}
	}
}
