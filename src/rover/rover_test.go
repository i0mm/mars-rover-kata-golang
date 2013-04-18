package rover

import "testing"


func TestCommandString(t *testing.T) {
  r := Rover{}
  r.Init(0, 0, 0)

  r.Command("ffrff")
  if(r.CurrentPosition().Point.X != 2 || r.CurrentPosition().Point.Y != 2 || r.CurrentPosition().Degree != 90) {
    t.Errorf("expected 2,2,90, got %v", r.CurrentPosition())
  } 
}