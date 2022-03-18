package genstack

import "testing"

func checkStackLen[T any](t *testing.T, s *Stack[T], len int) bool {
	if n := s.Len(); n != len {
		t.Errorf("s.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func TestStack(t *testing.T) {
	s := New[bool]()
	checkStackLen[bool](t, s, 0)
}

func TestPush(t *testing.T) {
	s := New[bool]()
	checkStackLen[bool](t, s, 0)
	s.Push(true)
	checkStackLen[bool](t, s, 1)
	s.Push(false)
	checkStackLen[bool](t, s, 2)
}

func TestPop(t *testing.T) {
	s := New[bool]()
	s.Push(true)
	s.Push(false)
	s.Push(true)
	s.Push(false)
	checkStackLen[bool](t, s, 4)
	e := s.Pop()
	checkStackLen[bool](t, s, 3)
	if e.Value != false {
		t.Errorf("e.Value = %v, want %v", e.Value, false)
	}
	e = s.Pop()
	checkStackLen[bool](t, s, 2)
	if e.Value != true {
		t.Errorf("e.Value = %v, want %v", e.Value, true)
	}
	e = s.Pop()
	checkStackLen[bool](t, s, 1)
	if e.Value != false {
		t.Errorf("e.Value = %v, want %v", e.Value, false)
	}
	e = s.Pop()
	checkStackLen[bool](t, s, 0)
	if e.Value != true {
		t.Errorf("e.Value = %v, want %v", e.Value, true)
	}
}
