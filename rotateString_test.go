package rotateString

import "testing"

func TestRotate(t *testing.T) {
	s := Rotate("Cesar")

	if len(s) == 0 {
		t.Errorf("Rotate return an empty string")
	}

	if s != "raseC" {
		t.Errorf("Rotate string fail, returned \"%s\"", s)
	}

}
