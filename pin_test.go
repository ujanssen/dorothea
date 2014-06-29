package dorothea_test

import (
	"github.com/ujanssen/dorothea"
	"testing"
)

func TestNewPin(t *testing.T) {
	pins := dorothea.NewPins(nil)
	want := 4
	if got := len(pins); got != want {
		t.Errorf("len(pins) = %v, want %v", got, want)
	}
}
