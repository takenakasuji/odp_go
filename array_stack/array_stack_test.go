package array_stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	if r := New().n; r != 0 {
		t.Errorf("ArrayStack.New().n = %v", r)
	}
}

func TestSet(t *testing.T) {
	tests := []int{100}

	as := New()
	as.buf = tests
	r := as.Set(0, 1)
	if r != 1 {
		t.Errorf("ArrayStack.Set() = %v", r)
	}
}