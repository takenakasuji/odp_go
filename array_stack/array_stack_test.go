package array_stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	if r := New().n; r != 0 {
		t.Errorf("ArrayStack.New().n = %v", r)
	}
}