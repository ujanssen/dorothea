package dorothea

type Pin struct {
	player *Player
	field  int
}

const OUT_AREA = -1

func NewPins(p *Player) [4]*Pin {
	var pins [4]*Pin
	for i := 0; i < 4; i++ {
		pins[i] = &Pin{
			player: p,
			field:  OUT_AREA}
	}
	return pins
}
