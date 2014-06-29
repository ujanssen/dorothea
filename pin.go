package dorothea

type Position int

const (
	Out     Position = iota // Pin in OutArea
	OnField                 // Pin on Field
	InHome                  // Pin in HomeArea
)

type Pin struct {
	player    *Player
	position  Position
	iField    int // index in game.fields when position == OnField
	iHomeArea int // index in player.HomeArea when position == InHome
}

const NOT_IN_INDEX = -1

func NewPins(p *Player) [4]*Pin {
	var pins [4]*Pin
	for i := 0; i < 4; i++ {
		pins[i] = &Pin{
			player:    p,
			position:  Out,
			iField:    NOT_IN_INDEX,
			iHomeArea: NOT_IN_INDEX}
	}
	return pins
}
