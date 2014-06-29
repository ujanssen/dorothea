package dorothea

import "math/rand"

var StartField map[Color]int = map[Color]int{Yellow: 0, Red: 10, Blue: 20, Green: 30}

type Player struct {
	game  *Game
	Color Color
	pins  [4]*Pin
}

func NewPlayer(g *Game, c Color) *Player {
	p := &Player{
		game:  g,
		Color: c}

	p.pins = NewPins(p)
	return p
}

func NewPlayers(g *Game, playerColors []Color) []*Player {
	players := make([]*Player, len(playerColors))
	for i, _ := range players {
		players[i] = NewPlayer(g, playerColors[i])
	}
	return players
}

func (p *Player) ThrowDice() int {
	return rand.Intn(6) + 1
}

// A player has won, if he has four pieces in the home area
func (p *Player) HasWon() bool {
	return p.PinsInHomeArea() == 4
}

func (p *Player) PinsInHomeArea() int {
	number := 0
	for _, pin := range p.pins {
		if pin.IsInHomeArea() {
			number++
		}
	}
	return number
}
func (p *Player) PinsInOutArea() int {
	number := 0
	for _, pin := range p.pins {
		if pin.IsInOutArea() {
			number++
		}
	}
	return number
}

func (p *Player) PinsOnField() int {
	number := 0
	for _, pin := range p.pins {
		if pin.IsOnField() {
			number++
		}
	}
	return number
}

func (p *Player) GetPinOnField(field int) *Pin {
	for _, pin := range p.pins {
		if pin.IsOnField() && pin.FieldIndex() == field {
			return pin
		}
	}
	return nil
}

/* unused
func (p *Player) PlayMove(g *Game) {
	dice := p.ThrowDice()
	p.PlayMoveWithDice(g, dice)

}
*/

func (p *Player) Pin(no int) *Pin {
	return p.pins[no]
}

func (p *Player) StartField() int {
	return StartField[p.Color]
}

func (p *Player) IsOwnPieceOnField(field int) bool {
	return p.game.GetColorOnField(field) == p.Color
}

func (p *Player) CanMoveTo(field int) bool {
	return !p.IsOwnPieceOnField(field)
}

func (p *Player) PlayMoveWithDice(dice int) {
	start := StartField[p.Color]

	if dice == 6 && p.PinsInOutArea() > 0 && p.CanMoveTo(start) {
		p.game.SetPlayerOnField(p, start)
		p.MovePinFromOutAreaToField(start)
	}
}

func (p *Player) MovePinFromOutAreaToField(field int) {
	for _, pin := range p.pins {
		if pin.IsInOutArea() {
			pin.MovePinFromOutAreaToField(field)
			return
		}
	}
}
func (p *Player) MovePinToHomeArea(iPin, iHome int) {
	p.pins[iPin].MoveToHomeArea(iHome)
}
