package dorothea

import "math/rand"

type Player struct {
	Color   Color
	OutArea int
}

func createPlayers(playerColors []Color) []Player {
	players := make([]Player, len(playerColors))
	for i, _ := range players {
		players[i] = Player{OutArea: 4, Color: playerColors[i]}
	}
	return players
}

func (p Player) ThrowDice() int {
	return rand.Intn(6) + 1
}
