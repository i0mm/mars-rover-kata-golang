package rover

type Rover struct {
  Path []position
}

func (r *Rover) CurrentPosition() *position {
  if len(r.Path) > 0 {
    return &r.Path[len(r.Path)-1]
  }
  return nil
}

func (r *Rover) Place(x, y int, degree int16) *Rover {
  r.Path = append(r.Path, position{point{x, y}, degree})
  return r
} 


func (r *Rover) OneLeftTurn() *Rover {
  r.Turn(-90)
  return r
}

func (r *Rover) OneRightTurn() *Rover {
  r.Turn(90)
  return r
}

func (r *Rover) Forward() *Rover {
  r.Move(1)
  return r
}

func (r *Rover) Backward() *Rover {
  r.Move(-1)
  return r
}


func (r *Rover) Move(ticks int) *Rover {
  d := *r.CurrentPosition()
  d.MoveInDirection(ticks)
  r.Path = append(r.Path, d)
  return r
}

func (r *Rover) Turn(degree int16) *Rover {
  d := *r.CurrentPosition()
  d.Rotate(degree);
  r.Path = append(r.Path, d)
  return r
}