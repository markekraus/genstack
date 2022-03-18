// Package genstack provides a generic FILO/LIFO stack.
package genstack

// Element is a element from the stack
type Element[T any] struct {
	next, prev *Element[T]
	stack      *Stack[T]
	Value      T
}

type Stack[T any] struct {
	root Element[T]
	len  int
}

// lazyInit lazily initializes a zero stack value.
func (s *Stack[T]) lazyInit() {
	if s.root.next == nil {
		s.Init()
	}
}

// insert inserts e after at, increments s.len, and returns e.
func (s *Stack[T]) insert(e, at *Element[T]) *Element[T] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.stack = s
	s.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (s *Stack[T]) insertValue(v T, at *Element[T]) *Element[T] {
	return s.insert(&Element[T]{Value: v}, at)
}

// remove removes e from its stack, decrements s.len
func (s *Stack[T]) remove(e *Element[T]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.stack = nil
	s.len--
}

// Push adds a new Element[T] e with value v at the top of Stack[T] s and returns e.
func (s *Stack[T]) Push(v T) *Element[T] {
	s.lazyInit()
	return s.insertValue(v, s.root.prev)
}

// Pop returns the next Element[T] e and removes it from Stack[T] s or nil if the stack is empty.
func (s *Stack[T]) Pop() *Element[T] {
	if s.len == 0 {
		return nil
	}
	e := s.root.prev
	s.remove(e)
	return e
}

// Len returns the number of elements of Stack[T] s.
// The complexity is O(1).
func (s *Stack[T]) Len() int { return s.len }

// Init initializes or clears stack s.
func (s *Stack[T]) Init() *Stack[T] {
	s.root.next = &s.root
	s.root.prev = &s.root
	s.len = 0
	return s
}

// New returns an initialized stack.
func New[T any]() *Stack[T] {
	return new(Stack[T]).Init()
}
