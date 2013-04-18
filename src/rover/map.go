package rover

type Map struct {
	height    int
	width     int
	obstacles []obstacle
	rovers    []Rover
}
