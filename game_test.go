package dorothea_test

import (
	"github.com/ujanssen/dorothea"
	"testing"
)

// go test src/github.com/ujanssen/dorothea/game_test.go

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

// Each player has 4 game pieces, which are in the "out" area when the game starts
func TestGameOutAreaAtStart(t *testing.T) {
	g, _ := dorothea.NewGame(getNPlayerColors(4))
	for _, player := range g.Players {
		if player.OutArea != 4 {
			t.Errorf("player.OutArea = %v, want 4", player.OutArea)
		}
	}
}

// Start a Game with Red and Green.
// Red should be the first active player.
// Green should be the next player.
// Red should be the next player.
func TestGameFirstPlayerAtStart(t *testing.T) {
	playerColors := []dorothea.Color{dorothea.Red, dorothea.Green}
	g, _ := dorothea.NewGame(playerColors)

	c := g.CurrentPlayer().Color
	if c != dorothea.Red {
		t.Errorf("g.CurrentPlayer().Color = %v, want dorothea.Red", c)
	}

	e := g.NextPlayer().Color
	if e != dorothea.Green {
		t.Errorf("g.NextPlayer().Color = %v, want dorothea.Green", e)
	}

	r := g.NextPlayer().Color
	if r != dorothea.Red {
		t.Errorf("g.NextPlayer().Color = %v, want dorothea.Red", r)
	}
}

// At game start, all fields should be empty
func TestGameFieldsAtStart(t *testing.T) {
	playerColors := []dorothea.Color{dorothea.Red, dorothea.Green}
	g, _ := dorothea.NewGame(playerColors)

	for field := 0; field < 40; field++ {
		if g.GetColorOnField(field) != dorothea.Empty {
			t.Errorf("field = %v, want dorothea.Empty", field)

		}
	}
}

func TestGameOver(t *testing.T) {
	g := newGame()
	p := g.CurrentPlayer()

	p.HomeArea[1] = true
	p.HomeArea[2] = true
	p.HomeArea[3] = true

	if g.GameOver() == true {
		t.Errorf("p.GameOver() = true, want false")
	}

	p.HomeArea[0] = true

	if g.GameOver() == false {
		t.Errorf("p.GameOver() = false, want true")
	}
}

// private ----------------------------------------------------

func newGame() *dorothea.Game {
	g, _ := dorothea.NewGame([]dorothea.Color{dorothea.Red, dorothea.Green})
	return g
}

func getNPlayerColors(n int) []dorothea.Color {
	playerColors := make([]dorothea.Color, 0)
	if n > 1 {
		playerColors = append(playerColors, dorothea.Yellow)
		playerColors = append(playerColors, dorothea.Red)
	}
	if n > 2 {
		playerColors = append(playerColors, dorothea.Blue)
	}
	if n > 3 {
		playerColors = append(playerColors, dorothea.Green)
	}
	if n > 4 {
		playerColors = append(playerColors, 8)
	}
	return playerColors
}

// Test with valid n players
func testGameWithValidNPlayers(n, e int, t *testing.T) {
	g, err := dorothea.NewGame(getNPlayerColors(n))
	if err != nil {
		t.Errorf("dorothea.NewGame(%v), got error %v, want no error", n, g.Players, e)
	}
	if e != len(g.Players) {
		t.Errorf("dorothea.NewGame(%v) = %v, want %v", n, g.Players, e)
	}
}

// Test with invalid n players
func testGameWithInvalidNPlayers(n int, t *testing.T) {
	g, err := dorothea.NewGame(getNPlayerColors(n))
	if err == nil {
		t.Errorf("dorothea.NewGame(%v) = %v, want error", n, g.Players)
	}
}
