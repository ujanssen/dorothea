package dorothea

import (
	"errors"
)

type Game struct {
	Players      []*Player
	fields       []Color
	activePlayer int
}

func NewGame(playerColors []Color) (*Game, error) {
	p := len(playerColors)

	if p > 1 && p <= 4 {
		g := &Game{
			fields:       make([]Color, 40),
			activePlayer: 0}

		(*g).Players = NewPlayers(g, playerColors)

		return g, nil
	}
	return &Game{}, errors.New("The game can be played by 2, 3 or 4 players")
}

func (g *Game) CurrentPlayer() *Player {
	return g.Players[g.activePlayer]
}

func (g *Game) NextPlayer() *Player {
	g.activePlayer = g.activePlayer + 1
	if g.activePlayer == len(g.Players) {
		g.activePlayer = 0
	}
	return g.CurrentPlayer()
}

func (g *Game) GameOver() bool {
	return g.CurrentPlayer().HasWon()
}

func (g *Game) SetPlayerOnField(p *Player, field int) {
	g.fields[field] = p.Color
}
func (g *Game) GetColorOnField(field int) Color {
	return g.fields[field]
}

/* unused
func (g *Game) PlayTheGame() {
	for {
		g.CurrentPlayer().PlayMove(g)
		if g.GameOver() {
			return
		}
		g.NextPlayer()
	}
}
*/
