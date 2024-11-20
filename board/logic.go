package board

// PlayerWins erwartet ein Zeichen.
// Gibt `true` zurück, wenn es eine Zeile, Spalte oder Diagonale gibt, die nur aus dem Zeichen besteht.
func (b Board) PlayerWins(s string) bool {
	for i := 0; i < len(b); i++ {
		if b.RowContainsOnly(i, s) || b.ColContainsOnly(i, s) {
			return true
		}
	}
	return b.DiagContainsOnly(0, s) || b.DiagContainsOnly(1, s)
}

// HINT
// Nutzen Sie eine Schleife, um über die Zeilen und Spalten zu iterieren.
// Falls Sie eine Zeile oder Spalte finden, die nur aus dem Zeichen besteht, geben Sie vorzeitig `true` zurück.

// GameOver gibt `true` zurück, wenn das Spiel vorbei ist.
// Das Spiel ist vorbei, wenn ein Spieler gewonnen hat oder das Board voll ist.
func (b Board) GameOver() bool {
	return b.Full() || b.PlayerWins("X") || b.PlayerWins("O")
}

// HINT
// Nutzen Sie die Funktionen `Full` und `PlayerWins`.
