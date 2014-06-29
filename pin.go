package dorothea

type Pin struct {
	player *Player
	field  int
}

func NewPin(p *Player) *Pin {
	pin := &Pin{player: p}

	return pin
}
