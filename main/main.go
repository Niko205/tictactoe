package main

import (
	"fmt"
	board "tictactoe/board_done"
)

func main() {
	s := ""
	var x, y int
	for {
		b := board.New(3, 3, " ")
		fmt.Println(board.Board.String(b))
		for {
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
			fmt.Println(board.Board.String(b))
			if b.GameOver() {
				break
			}
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
			fmt.Println(board.Board.String(b))
			if b.GameOver() {
				break
			}
		}
		fmt.Println("Neues Spiel? J/N")
		fmt.Scan(&s)
		if s == "N" {
			break
		}
	}
}
