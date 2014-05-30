package dorothea

import (
	"errors"
)

type GamePlay interface {
	newGame(players int) Game
}

type Game struct {
	Players int
}

func NewGame(players int) (Game, error) {
	if players > 1 && players <= 4 {
		g := Game{Players: players}
		return g, nil

	}
	return Game{}, errors.New("The game can be played by 2, 3 or 4 players")
}
