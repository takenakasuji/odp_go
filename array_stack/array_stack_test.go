package array_stack

import (
	"reflect"
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

func TestGet(t *testing.T) {
	tests := []int{100, 1, 20}

	as := New()
	as.buf = tests
	as.Set(2, 200)
	r := as.Get(2)
	if r != 200 {
		t.Errorf("ArrayStack.Get() = %v", r)
	}
}

func TestAdd(t *testing.T) {
	tests := []int{100, 1, 20}

	as := New()
	as.buf = append([]int{}, tests...)
	as.Add(2, 200)
	if reflect.DeepEqual(tests, as.buf) {
		t.Errorf("ArrayStack.Add() = %v, %v", tests, as.buf)
	}
}
