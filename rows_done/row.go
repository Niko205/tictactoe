package rows

type Row []string

// New erwartet eine Länge und einen String.
// Gibt eine neue `Row` zurück, die mit dem String gefüllt ist.
func New(length int, fill string) Row {
	r := make(Row, length)
	for i := 0; i < length; i++ {
		r[i] = fill
	}

	return r
}

// String gibt die `Row` als String zurück.
// Die Elemente sind durch ` | ` getrennt und von `|` umgeben.
func (r Row) String() string {
	s := "|"
	for i := 0; i < len(r); i++ {
		s += " " + r[i] + " " + "|"
	}
	return s
}

// ContainsOnly erwartet einen String.
// Gibt `true` zurück, wenn die `Row` nur aus dem String besteht.
func (r Row) ContainsOnly(s string) bool {
	for _, n := range r {
		if n != s {
			return false
		}
	}
	return true
}
