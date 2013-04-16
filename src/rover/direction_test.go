package rover

import (
  "testing"
)

func TestRotate(t *testing.T) {
  d := direction{}

  d = direction{point{0,0}, 0}
  if d2 := d.Rotate(90); d2.Degree != 90 {
    t.Errorf("expected 90, got %v", d2.Degree)
  }

  d = direction{point{0,0}, 0}
  if d2 := d.Rotate(-90); d2.Degree != -90 {
    t.Errorf("expected -90, got %v", d2.Degree)
  }

  d = direction{point{0,0}, 0}
  if d.Rotate(45); d.Degree != 45 {
    t.Errorf("expected 45, got %v", d.Degree)
  }

  d = direction{point{0,0}, 0}
  if d.Rotate(-45); d.Degree != -45 {
    t.Errorf("expected -45, got %v", d.Degree)
  }


  d = direction{point{0,0}, 0}
  if d.Rotate(360); d.Degree != 0 {
    t.Errorf("expected 0, got %v", d.Degree)
  }

  d = direction{point{0,0}, 0}
  if d.Rotate(-360); d.Degree != 0 {
    t.Errorf("expected 0, got %v", d.Degree)
  }

  d = direction{point{0,0}, 0}
  if d.Rotate(361); d.Degree != 1 {
    t.Errorf("expected 1, got %v", d.Degree)
  }

  d = direction{point{0,0}, 0}
  if d.Rotate(-361); d.Degree != -1 {
    t.Errorf("expected -1, got %v", d.Degree)
  }


}
