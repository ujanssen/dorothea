package dorothea_test

import (
	"github.com/ujanssen/dorothea"
	"testing"
)

// go test src/github.com/ujanssen/dorothea/player_test.go

// The dice has number from 1 to 6
func TestPlayerThrowDice(t *testing.T) {
	g := newGame()
	n := 1
	p := dorothea.NewPlayer(g, dorothea.Green)
	for n < 1000 {
		number := p.ThrowDice()
		if number < 1 || number > 6 {
			t.Errorf("p.TrowDice() = %v, want number 1,2,3,4,5 or 6", number)
		}
		n = n + 1
	}
}

func TestPlayerHasWon(t *testing.T) {
	g := newGame()
	p := g.CurrentPlayer()

	p.HomeArea[1] = p.Pin(1)
	p.HomeArea[2] = p.Pin(2)
	p.HomeArea[3] = p.Pin(3)

	if p.HasWon() == true {
		t.Errorf("p.HasWon() = true, want false")
	}

	p.HomeArea[0] = p.Pin(0)

	if p.HasWon() == false {
		t.Errorf("p.HasWon() = false, want true")
	}
}

func TestPlayerOnStartThrowsASix(t *testing.T) {
	g := newGame()
	p := g.CurrentPlayer()
	start := p.StartField()

	p.PlayMoveWithDice(6)

	if p.OutArea != 3 {
		t.Errorf("p.OutArea = %v, want 3", p.OutArea)
	}
	if g.GetColorOnField(start) != p.Color {
		t.Errorf("g.Fields[%v] = %v, want %v", start, g.GetColorOnField(start), p.Color)
	}
	if len(p.Position) != 1 {
		t.Errorf("len(p.Position) = %v, want 1", len(p.Position))
	}
	if p.Position[0] != start {
		t.Errorf("p.Position[0] = %v, want %v", p.Position[0], start)
	}
}
