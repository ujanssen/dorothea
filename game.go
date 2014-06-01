package dorothea

import (
	"errors"
)

type Game struct {
	Players      []*Player
	Fields       []Color
	activePlayer int
}

func NewGame(playerColors []Color) (*Game, error) {
	p := len(playerColors)

	if p > 1 && p <= 4 {
		players := createPlayers(playerColors)
		g := &Game{Players: players, Fields: make([]Color, 40), activePlayer: 0}
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
