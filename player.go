package dorothea

import "math/rand"

var StartField map[Color]int = map[Color]int{Yellow: 0, Red: 10, Blue: 20, Green: 30}

type Player struct {
	game     *Game
	Color    Color
	OutArea  int
	Position []int  // of pieces on board
	HomeArea []bool // piece in home[0,1,2,3] ?
}

func NewPlayer(g *Game, c Color) *Player {
	p := Player{
		game:     g,
		OutArea:  4,
		Color:    c,
		Position: make([]int, 0),
		HomeArea: make([]bool, 4)}
	return &p
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
		if piece {
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
