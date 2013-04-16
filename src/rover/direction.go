package rover

type direction struct {
  Point point
  Degree int
}

func (d *direction) Rotate(degree int) *direction {
  d.Degree = degree % 360
  return d // keep it chainable
}