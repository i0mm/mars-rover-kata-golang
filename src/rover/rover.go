package rover

type Rover struct {
  Path []direction
}

func (r *Rover) CurrentDirection() *direction {
  if len(r.Path) > 0 {
    return &r.Path[len(r.Path)-1]
  }
  return nil
}

func (r *Rover) Place(x, y, degree int) *Rover {
  r.Path = append(r.Path, direction{point{x, y}, degree})
  return r
} 

func (r *Rover) Forward() *Rover {
  d := *r.CurrentDirection()
  d.Point.X++
  r.Path = append(r.Path, d)
  return r
}