package dorothea_test

import (
	"github.com/ujanssen/dorothea"
	"testing"
)

func TestNewPin(t *testing.T) {
	pin := dorothea.NewPin(nil)
	want := true
	if got := pin.IsInOutArea(); got != want {
		t.Errorf("pin.IsInOutArea() = %v, want %v", got, want)
	}
}

func TestNewPins(t *testing.T) {
	pins := dorothea.NewPins(nil)
	want := 4
	if got := len(pins); got != want {
		t.Errorf("len(pins) = %v, want %v", got, want)
	}
}

func TestMoveTo(t *testing.T) {
	pin := dorothea.NewPin(nil)
	pin.MoveTo(33)
	want := true
	if got := pin.IsOnField(); got != want {
		t.Errorf("pin.IsOnField() = %v, want %v", got, want)
	}
}
