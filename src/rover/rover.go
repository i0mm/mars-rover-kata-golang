package rover


func Move(x int) point {
  p := direction{point{1, 2}, 90}
  p.Point.X = 2
  return p.Point
}