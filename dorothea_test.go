package dorothea

import (
	"github.com/ujanssen/dorothea"
	"testing"
)

// go test src/github.com/ujanssen/dorothea/dorothea_test.go

// The with valid n players
func testGameWithValidNPlayers(n, e int, t *testing.T) {
	g, _ := dorothea.NewGame(n)
	if e != g.Players {
		t.Errorf("dorothea.NewGame(%v) = %v, want %v", n, g.Players, e)
	}
}

// The with invalid n players
func testGameWithInvalidNPlayers(n int, t *testing.T) {
	g, err := dorothea.NewGame(n)
	if err == nil {
		t.Errorf("dorothea.NewGame(%v) = %v, want error", n, g.Players)
	}
}

// The game can be played by 2, 3 or 4 players
func TestGameWith2Players(t *testing.T) {
	testGameWithValidNPlayers(2, 2, t)
}
func TestGameWith3Players(t *testing.T) {
	testGameWithValidNPlayers(3, 3, t)
}
func TestGameWith4Players(t *testing.T) {
	testGameWithValidNPlayers(4, 4, t)
}

// The game can not be played by 1 or 5 players
func TestGameWith1Player(t *testing.T) {
	testGameWithInvalidNPlayers(1, t)
}
func TestGameWith5Players(t *testing.T) {
	testGameWithInvalidNPlayers(5, t)
}
