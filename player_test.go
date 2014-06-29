package dorothea_test

import (
	"github.com/ujanssen/dorothea"
	"testing"
)

func TestPiecesInOutAreaOnStart(t *testing.T) {
	g := newGame()
	p := g.CurrentPlayer()

	want := 4
	if got := p.PiecesInOutArea(); got != want {
		t.Errorf("p.PiecesInOutArea() = %v, want %v", got, want)
	}

}

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
		n++
	}
}

func TestPlayerHasWon(t *testing.T) {
	g := newGame()
	p := g.CurrentPlayer()

	p.MovePinToHomeArea(1, 1)
	p.MovePinToHomeArea(2, 2)
	p.MovePinToHomeArea(3, 3)

	if p.HasWon() == true {
		t.Errorf("p.HasWon() = true, want false")
	}

	p.MovePinToHomeArea(0, 0)

	if p.HasWon() == false {
		t.Errorf("p.HasWon() = false, want true")
	}
}

func TestPlayerOnStartThrowsASix(t *testing.T) {
	g := newGame()
	p := g.CurrentPlayer()
	start := p.StartField()

	p.PlayMoveWithDice(6)

	if p.PiecesInOutArea() != 3 {
		t.Errorf("p.OutArea = %v, want 3", p.PiecesInOutArea())
	}
	if g.GetColorOnField(start) != p.Color {
		t.Errorf("g.Fields[%v] = %v, want %v", start, g.GetColorOnField(start), p.Color)
	}
	if p.PiecesOnField() != 1 {
		t.Errorf("p.PiecesOnField() = %v, want 1", p.PiecesOnField())
	}
	if p.GetPinOnField(start) == nil {
		t.Errorf("p.GetPinOnField(start) = %v, want %v", p.GetPinOnField(start), start)
	}
}
