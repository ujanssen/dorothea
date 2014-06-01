package dorothea

import (
	"github.com/ujanssen/dorothea"
	"testing"
)

// go test src/github.com/ujanssen/dorothea/player_test.go

// The dice has number from 1 to 6
func TestPlayerThrowDice(t *testing.T) {
	n := 1
	p := dorothea.NewPlayer(dorothea.Green)
	for n < 1000 {
		number := p.ThrowDice()
		if number < 1 || number > 6 {
			t.Errorf("p.TrowDice() = %v, want number < 1 && number > 6", number)
		}
		n = n + 1
	}
}

func TestPlayerHasWon(t *testing.T) {
	p := dorothea.NewPlayer(dorothea.Green)

	p.HomeArea[1] = true
	p.HomeArea[2] = true
	p.HomeArea[3] = true

	if p.HasWon() == true {
		t.Errorf("p.HasWon() = true, want false")
	}

	p.HomeArea[0] = true

	if p.HasWon() == false {
		t.Errorf("p.HasWon() = false, want true")
	}
}

func TestPlayerOnStartThrowsASix(t *testing.T) {
	g, _ := dorothea.NewGame([]dorothea.Color{dorothea.Red, dorothea.Green})

	p := g.CurrentPlayer()
	start := p.StartField()

	p.PlayMoveWithDice(g, 6)

	if p.OutArea != 3 {
		t.Errorf("p.OutArea = %v, want 3", p.OutArea)
	}
	if g.Fields[start] != p.Color {
		t.Errorf("g.Fields[%v] = %v, want %v", start, g.Fields[start], p.Color)
	}
	if len(p.Position) != 1 {
		t.Errorf("len(p.Position) = %v, want 1", len(p.Position))
	}
	if p.Position[0] != start {
		t.Errorf("p.Position[0] = %v, want %v", p.Position[0], start)
	}
}
