package rover


func Move(x int) int {
  p := direction{point{1, 2}, 90}
  p.Point.X = 4
  return p.Point.X
}