package dorothea

import "math/rand"

var StartField map[Color]int = map[Color]int{Yellow: 0, Red: 10, Blue: 20, Green: 30}

type Player struct {
	game     *Game
	Color    Color
	OutArea  int
	Position []int   // of pieces on board
	HomeArea [4]*Pin // pin in home[0,1,2,3] ?
	pins     [4]*Pin
}

func NewPlayer(g *Game, c Color) *Player {
	var homeArea [4]*Pin
	p := &Player{
		game:     g,
		OutArea:  4,
		Color:    c,
		Position: make([]int, 0),
		HomeArea: homeArea}

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
	return p.PiecesInHomeArea() == 4
}

func (p *Player) PiecesInHomeArea() int {
	number := 0
	for _, piece := range p.HomeArea {
		if piece != nil {
			number = number + 1
		}
	}
	return number
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

	if dice == 6 && p.OutArea > 0 && p.CanMoveTo(start) {
		p.game.SetPlayerOnField(p, start)
		p.OutArea = p.OutArea - 1
		p.Position = append(p.Position, start)
	}
}
