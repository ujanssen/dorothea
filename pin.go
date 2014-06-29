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
	iHomeArea int // 0,1,2,3
}

const NOT_IN_INDEX = -1

func NewPin(p *Player) *Pin {
	return &Pin{
		player:    p,
		position:  Out,
		iField:    NOT_IN_INDEX,
		iHomeArea: NOT_IN_INDEX}
}

func NewPins(p *Player) [4]*Pin {
	var pins [4]*Pin
	for i := 0; i < 4; i++ {
		pins[i] = NewPin(p)
	}
	return pins
}

func (pin *Pin) MoveTo(index int) {
	pin.position = OnField
	pin.iField = index
}

func (pin *Pin) MoveToHomeArea(index int) {
	pin.position = InHome
	pin.iHomeArea = index
}

func (pin *Pin) IsInOutArea() bool {
	return pin.position == Out
}
func (pin *Pin) IsOnField() bool {
	return pin.position == OnField
}
func (pin *Pin) IsInHomeArea() bool {
	return pin.position == InHome
}
func (pin *Pin) FieldIndex() int {
	return pin.iField
}

func (pin *Pin) MovePinFromOutAreaToField(field int) {
	if !pin.IsInOutArea() {
		panic("!pin.IsInOutArea()")
	}
	pin.MoveTo(field)
}
