package rotateString

import "testing"

func TestRotate(t *testing.T) {
	s := Rotate("Cesar")

	if s != "raseC" {
		t.Errorf("Rotate string fail")
	}
}
