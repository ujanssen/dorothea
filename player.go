package dorothea

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
