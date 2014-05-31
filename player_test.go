package dorothea

import (
	"github.com/ujanssen/dorothea"
	"testing"
)

// go test src/github.com/ujanssen/dorothea/player_test.go

// The dice has number from 1 to 6
func TestPlayerThrowDice(t *testing.T) {
	n := 1
	p := dorothea.Player{}
	for n < 1000 {
		number := p.ThrowDice()
		if number < 1 || number > 6 {
			t.Errorf("p.TrowDice() = %v, want number < 1 && number > 6", number)
		}
		n = n + 1
	}
}
