package rows

import (
	"fmt"
	"strings"
)

type Row []string

// New erwartet eine Länge und einen String.
// Gibt eine neue `Row` zurück, die mit dem String gefüllt ist.
func New(length int, fill string) Row {
	// SOLUTION
	return strings.Split(strings.Repeat(fill, length), "")
	// SOLUTION_END
}

// HINT
// Nutzen Sie die Funktionen `strings.Split` und `strings.Repeat`.

// String gibt die `Row` als String zurück.
// Die Elemente sind durch ` | ` getrennt und von `|` umgeben.
func (r Row) String() string {
	// SOLUTION
	return fmt.Sprintf("| %v |", strings.Join(r, " | "))
	// SOLUTION_END
}

// HINT
// Nutzen Sie die Funktionen `fmt.Sprintf` und `strings.Join`.

// ContainsOnly erwartet einen String.
// Gibt `true` zurück, wenn die `Row` nur aus dem String besteht.
func (r Row) ContainsOnly(s string) bool {
	// SOLUTION
	for _, v := range r {
		if v != s {
			return false
		}
	}
	// SOLUTION_END
	return true
}

// HINT
// Nutzen Sie eine Schleifen, um durch die Row zu iterieren.
// Falls Sie ein Element finden, das nicht s entspricht, geben Sie vorzeitig `false` zurück.
