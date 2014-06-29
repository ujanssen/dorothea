package dorothea_test

import (
	"github.com/ujanssen/dorothea"
	"testing"
)

func TestNewPin(t *testing.T) {
	pin := dorothea.NewPin(nil)
	if pin == nil {
		t.Errorf("pin == nil")
	}
}
