package array_stack

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	if r := len(New().buf); r != 0 {
		t.Errorf("ArrayStack.New().buf = %v", r)
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
	tests := make([]int, 0, 0)
	tests = append(tests, 1, 100, 20)

	as := New()
	as.buf = append([]int{}, tests...)
	p := 1
	as.Add(p, 55)
	if reflect.DeepEqual(tests, as.buf) {
		t.Errorf("ArrayStack.Add() = %v, %v", tests, as.buf)
	}
	if as.buf[p] != 55 {
		t.Errorf("ArrayStack.Add() invalid result: %v", as.buf[p])
	}
}

func TestRemove(t *testing.T) {
	tests := make([]int, 0, 0)
	tests = append(tests, 1, 100, 20)

	as := New()
	as.buf = append([]int{}, tests...)
	as.Remove(1)
	if reflect.DeepEqual(tests, as.buf) {
		t.Errorf("ArrayStack.Remove() = %v, %v", tests, as.buf)
	}
}

func TestCapacity(t *testing.T) {
	as := New()
	for i := 0; i <= 5; i++ {
		as.Add(i, i)
	}
	if cap(as.buf) != 8 {
		t.Errorf("ArrayStack.resize(): Incorrect proccesing")
	}

	for i := 0; i < 4; i++ {
		as.Remove(0)
	}
	if cap(as.buf) != 4 {
		t.Errorf("ArrayStack.resize(): Incorrect proccesing")
	}
}
