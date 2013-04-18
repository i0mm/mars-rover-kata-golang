package rover

import "testing"


func TestCommandString(t *testing.T) {
  r := Rover{}
  r.Init(0, 0, 0)

  r.Command("ffrff")
  if(r.CurrentPosition().Point.X != 2 || r.CurrentPosition().Point.Y != 2 || r.CurrentPosition().Degree != 90) {
    t.Errorf("expected , got %v", r.CurrentPosition())
  }

  
}


/*
func TestForward(t *testing.T) {
  r := Rover{}

  r.Place(0, 0, 90)
  if(len(r.Path) != 1) {
    t.Errorf("rover has 1 path item")
  }

  if r.Forward(); len(r.Path) != 3 && r.CurrentPosition().Point.X != 99 {
    t.Errorf("expected , got %v", r.CurrentPosition())
  }
}
*/