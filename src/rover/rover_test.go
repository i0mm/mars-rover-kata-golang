package rover

import "testing"

func TestCurrentDirection(t *testing.T) {
  r := Rover{}

  if(r.CurrentDirection() != nil) {
    t.Fail()
  }

  r.Place(1, 2, 90)
}


func TestForward(t *testing.T) {
  r := Rover{}

  r.Place(1, 2, 90)
  if(len(r.Path) != 1) {
    t.Errorf("rover has 1 path item")
  }

  if r.Forward(); len(r.Path) != 3 && r.CurrentDirection().Point.X != 99 {
    t.Errorf("expected , got %v", r.CurrentDirection())
  }
}