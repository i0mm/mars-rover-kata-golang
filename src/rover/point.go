package rover

type point struct {
  X, Y int
}

// add two points
func (p *point) Add(p2 point) *point {
  p.X += p2.X
  p.Y += p2.Y
  return p
}


