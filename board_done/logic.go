package board

// PlayerWins erwartet ein Zeichen.
// Gibt `true` zurück, wenn es eine Zeile, Spalte oder Diagonale gibt, die nur aus dem Zeichen besteht.
func (b Board) PlayerWins(s string) bool {
	for i := 0; i < len(b); i++ {
		if b.RowContainsOnly(i, s) {
			return true
		}
	}
	for i := 0; i < len(b); i++ {
		if b.ColContainsOnly(i, s) {
			return true
		}
	}
	for i := 0; i < 2; i++ {
		if b.DiagContainsOnly(i, s) {
			return true
		}
	}
	return false
}

// GameOver gibt `true` zurück, wenn das Spiel vorbei ist.
// Das Spiel ist vorbei, wenn ein Spieler gewonnen hat oder das Board voll ist.
func (b Board) GameOver() bool {
	if b.PlayerWins("X") {
		return true
	}
	if b.PlayerWins("O") {
		return true
	}
	if b.Full() {
		return true
	}
	return false
}
