package board

import (
	"strings"
	rows "tictactoe/rows_done"
)

type Board []rows.Row

// New erwartet Höhe, Breite und ein Zeichen.
// Liefert ein neues `Board` zurück, das mit dem Zeichen gefüllt ist.
func New(height, width int, fill string) Board {
	board := make(Board, height)
	for i := 0; i < height; i++ {
		board[i] = rows.New(width, fill)
	}
	return board
}

// String gibt das `Board` als String zurück.
// Die Zeilen sind durch Trenner der Form `+---+---+---+` getrennt.
func (b Board) String() string {
	rowStrings := make([]string, len(b))
	rowStrings[0] = "+---+---+---+\n"
	for i := 0; i < len(rowStrings); i++ {
		rowStrings[i] += b[i].String() + "\n+---+---+---+\n"
	}
	return strings.Join(rowStrings, "")
}

// Row erwartet eine Zeilennummer und liefert diese Zeile zurück.
func (b Board) Row(row int) rows.Row {
	return b[row]
}

// Set erwartet eine Spaltennummer und liefert diese Spalte zurück.
// Der Rückgabetype ist Row, da Zeilen und Spalten gleich sind.
func (b Board) Col(col int) rows.Row {
	c := make(rows.Row, len(b))
	for i := 0; i < len(b); i++ {
		c[i] = b[i][col]
	}
	return c
}

// Diag erwartet eine Diagonalennummer und liefert diese Diagonale zurück.
// Der Rückgabetype ist Row, da Diagonalen und Zeilen gleich sind.
// Die Diagonalennummer ist 0 für die Hauptdiagonale und 1 für die Nebendiagonale.
func (b Board) Diag(diag int) rows.Row {
	d := make(rows.Row, len(b))
	if diag == 0 {
		for i := 0; i < len(b); i++ {
			d[i] = b[i][i]
		}
	} else {
		z := len(b) - 1
		for i := 0; i < len(b); i++ {
			d[i] = b[i][z]
			z--
		}
	}
	return d
}

func (b Board) Empty(row, col int) bool {

	return b[row][col] == " " || b[row][col] == "*"
}

// Set erwartet eine Zeilen- und eine Spaltennummer und ein Zeichen.
// Setzt das Zeichen an die entsprechende Stelle.
func (b Board) Set(row, col int, fill string) {
	b[row][col] = fill
}

// Full gibt `true` zurück, wenn das Board voll ist.
func (b Board) Full() bool {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == " " || b[i][j] == "*" {
				return false
			}
		}
	}
	return true
}

// RowContainsOnly erwartet eine Zeilennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Zeile nur aus dem Zeichen besteht.
func (b Board) RowContainsOnly(row int, s string) bool {

	return b[row].ContainsOnly(s)
}

// ColContainsOnly erwartet eine Spaltennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Spalte nur aus dem Zeichen besteht.
func (b Board) ColContainsOnly(col int, s string) bool {
	// TODO
	return b.Col(col).ContainsOnly(s)
}

// DiagContainsOnly erwartet eine Diagonalennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Diagonale nur aus dem Zeichen besteht.
// Die Diagonalennummer ist 0 für die Hauptdiagonale und 1 für die Nebendiagonale.
func (b Board) DiagContainsOnly(diag int, s string) bool {
	return b.Diag(diag).ContainsOnly(s)
}
