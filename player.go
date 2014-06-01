package dorothea

import "math/rand"

type Player struct {
	Color    Color
	OutArea  int
	Position []int  // of pieces on board
	HomeArea []bool // piece in home[0,1,2,3] ?
}

func NewPlayer(c Color) *Player {
	p := Player{OutArea: 4, Color: c, Position: make([]int, 0), HomeArea: make([]bool, 4)}
	return &p
}

func createPlayers(playerColors []Color) []*Player {
	players := make([]*Player, len(playerColors))
	for i, _ := range players {
		players[i] = NewPlayer(playerColors[i])
	}
	return players
}

func (p Player) ThrowDice() int {
	return rand.Intn(6) + 1
}

// A player has won, if he has four pieces in the home area
func (p Player) HasWon() bool {
	return p.PiecesInHomeArea() == 4
}

func (p Player) PiecesInHomeArea() int {
	number := 0
	for _, piece := range p.HomeArea {
		if piece {
			number = number + 1
		}
	}
	return number
}
