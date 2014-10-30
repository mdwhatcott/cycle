package cycle

import "testing"
import "github.com/mdwhatcott/cycle"

func TestConst(t *testing.T) {
	if cycle.A != 42 {
		t.Error("Expected 42 but got:", cycle.A)
	}
}
