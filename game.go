package dorothea

import (
	"errors"
)

type Game struct {
	Players []Player
}

func NewGame(playerColors []Color) (Game, error) {
	p := len(playerColors)

	if p > 1 && p <= 4 {
		players := make([]Player, p)
		for i, _ := range players {
			players[i] = Player{OutArea: 4, Color: playerColors[i]}
		}
		g := Game{Players: players}
		return g, nil
	}
	return Game{}, errors.New("The game can be played by 2, 3 or 4 players")
}
