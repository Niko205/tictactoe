package board

import (
	"strings"
	"tictactoe/rows"
)

type Board []rows.Row

// New erwartet Höhe, Breite und ein Zeichen.
// Liefert ein neues `Board` zurück, das mit dem Zeichen gefüllt ist.
func New(height, width int, fill string) Board {
	board := make(Board, height)
	// TODO
	return board
}

// HINT
// Das mit `make` erstellte Board hat schon die richtige Länge, die Zeilen sind aber noch leer.
// Nutzen Sie eine Schleife, um die Zeilen zu füllen. Dabei können Sie die Funktion `rows.New` verwenden.

// String gibt das `Board` als String zurück.
// Die Zeilen sind durch Trenner der Form `+---+---+---+` getrennt.
func (b Board) String() string {
	rowStrings := make([]string, len(b))
	divider := "\n"
	// TODO
	return strings.Join(rowStrings, divider)
}

// HINT
// Definieren Sie zuerst den Trenner, der zwischen den Zeilen steht.
// Dazu können Sie die Funktion `strings.Repeat` verwenden.
// Iterieren Sie dann über die Zeilen und setzen für jede Zeile den Zeilenstring zusammen.
// Bei obiger  Vorgabe können sie diesen z.B. in das Array `rowStrings` speichern.
// Verwenden Sie am Ende `strings.Join`, um die Zeilen zu verbinden und `fmt.Sprintf`,
// um den Trenner noch am Anfang und am Ende hinzuzufügen.

// Row erwartet eine Zeilennummer und liefert diese Zeile zurück.
func (b Board) Row(row int) rows.Row {
	// TODO
	return rows.Row{}
}

// HINT
// Da die Zeilen bereits als `Row` definiert sind, können Sie die Zeile direkt zurückgeben.

// Set erwartet eine Spaltennummer und liefert diese Spalte zurück.
// Der Rückgabetype ist Row, da Zeilen und Spalten gleich sind.
func (b Board) Col(col int) rows.Row {
	c := make(rows.Row, len(b))
	// TODO
	return c
}

// HINT
// Die `make`-Vorgabe erstellt eine neue `Row` mit der richtigen Länge.
// Iterieren Sie über die Zeilen und fügen Sie das Element an der entsprechenden Stelle in die neue `Row` ein.

// Diag erwartet eine Diagonalennummer und liefert diese Diagonale zurück.
// Der Rückgabetype ist Row, da Diagonalen und Zeilen gleich sind.
// Die Diagonalennummer ist 0 für die Hauptdiagonale und 1 für die Nebendiagonale.
func (b Board) Diag(diag int) rows.Row {
	d := make(rows.Row, len(b))
	// TODO
	return d
}

// HINT
// Die `make`-Vorgabe erstellt eine neue `Row` mit der richtigen Länge.
// Iterieren Sie über das Board und fügen Sie die passenden Werte in d ein.
// Dabei müssen Sie unterscheiden, ob Sie die Haupt- oder Nebendiagonale betrachten.

// Set erwartet eine Zeilen- und eine Spaltennummer und ein Zeichen.
// Setzt das Zeichen an die entsprechende Stelle.
func (b Board) Set(row, col int, fill string) {
	// TODO
}

// HINT
// Da das Board effektig ein [][]String ist, können Sie das Zeichen direkt an der entsprechenden Stelle setzen.

// Full gibt `true` zurück, wenn das Board voll ist.
func (b Board) Full() bool {
	// TODO
	return true
}

// HINT
// Iterieren Süe in einer geschachtelten Schleife über das Board.
// Falls Sie ein leeres Feld finden, geben Sie vorzeitig `false` zurück.

// RowContainsOnly erwartet eine Zeilennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Zeile nur aus dem Zeichen besteht.
func (b Board) RowContainsOnly(row int, s string) bool {
	// TODO
	return false
}

// HINT
// Lassen Sie sich zuerst die Zeile berechnen und nutzen Sie dann die Funktion `ContainsOnly` von `Row`.

// ColContainsOnly erwartet eine Spaltennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Spalte nur aus dem Zeichen besteht.
func (b Board) ColContainsOnly(col int, s string) bool {
	// TODO
	return false
}

// HINT
// Lassen Sie sich zuerst die Spalte berechnen und nutzen Sie dann die Funktion `ContainsOnly` von `Row`.

// DiagContainsOnly erwartet eine Diagonalennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Diagonale nur aus dem Zeichen besteht.
// Die Diagonalennummer ist 0 für die Hauptdiagonale und 1 für die Nebendiagonale.
func (b Board) DiagContainsOnly(diag int, s string) bool {
	// TODO
	return false
}

// HINT
// Lassen Sie sich zuerst die Diagonale berechnen und nutzen Sie dann die Funktion `ContainsOnly` von `Row`.
