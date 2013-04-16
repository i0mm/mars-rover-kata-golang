package rover

import "testing"

func TestMove(t *testing.T) {
  if x := Move(2); x != 4 {
    t.Errorf("error %s", x)
  }
}