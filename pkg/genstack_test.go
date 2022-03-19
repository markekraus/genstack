package genstack

import "testing"

func checkStackLen[T any](t *testing.T, s *Stack[T], len int) bool {
	if n := s.Len(); n != len {
		t.Errorf("s.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

type myType struct {
	a, b int
}

func TestStack(t *testing.T) {
	s := New[bool]()
	checkStackLen(t, s, 0)
}

func TestPush(t *testing.T) {
	s := New[bool]()
	checkStackLen(t, s, 0)
	s.Push(true)
	checkStackLen(t, s, 1)
	s.Push(false)
	checkStackLen(t, s, 2)
}

func TestPop(t *testing.T) {
	s := New[bool]()
	s.Push(true)
	s.Push(false)
	s.Push(true)
	s.Push(false)
	checkStackLen(t, s, 4)
	e := s.Pop()
	checkStackLen(t, s, 3)
	if e.Value != false {
		t.Errorf("e.Value = %v, want %v", e.Value, false)
	}
	e = s.Pop()
	checkStackLen(t, s, 2)
	if e.Value != true {
		t.Errorf("e.Value = %v, want %v", e.Value, true)
	}
	e = s.Pop()
	checkStackLen(t, s, 1)
	if e.Value != false {
		t.Errorf("e.Value = %v, want %v", e.Value, false)
	}
	e = s.Pop()
	checkStackLen(t, s, 0)
	if e.Value != true {
		t.Errorf("e.Value = %v, want %v", e.Value, true)
	}
}

func TestStruct(t *testing.T) {
	e1 := &myType{1, 2}
	s := New[*myType]()
	s.Push(e1)
	checkStackLen(t, s, 1)
	e2 := s.Pop().Value
	if e2 != e1 {
		t.Errorf("e2 = %v, want %v", e2, e1)
	}
	if e2.a != 1 {
		t.Errorf("e2.a = %v, want %v", e2.a, 1)
	}
	if e2.b != 2 {
		t.Errorf("e2.b = %v, want %v", e2.b, 2)
	}
}
