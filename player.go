package dorothea

import "math/rand"

type Player struct {
	Color    Color
	OutArea  int
	Position []int  // of pieces on board
	HomeArea []bool // piece in home[0,1,2,3] ?
}

func createPlayers(playerColors []Color) []Player {
	players := make([]Player, len(playerColors))
	for i, _ := range players {
		players[i] = Player{OutArea: 4, Color: playerColors[i], Position: make([]int, 0), HomeArea: make([]bool, 4)}
	}
	return players
}

func (p Player) ThrowDice() int {
	return rand.Intn(6) + 1
}
