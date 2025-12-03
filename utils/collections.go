package utils

// Set is a simple set implementation using a map
type Set[T comparable] map[T]struct{}

// NewSet creates a new empty set
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Add adds an element to the set
func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

// Remove removes an element from the set
func (s Set[T]) Remove(item T) {
	delete(s, item)
}

// Contains checks if an element is in the set
func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

// Size returns the number of elements in the set
func (s Set[T]) Size() int {
	return len(s)
}

// ToSlice returns all elements as a slice
func (s Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for item := range s {
		result = append(result, item)
	}
	return result
}

// Queue is a simple FIFO queue
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Enqueue adds an item to the end of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first item from the queue
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Stack is a simple LIFO stack
type Stack[T any] struct {
	items []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

// Top returns the top item from the stack without removing it
func (s *Stack[T]) Top() T {
	return s.items[len(s.items)-1]
}

// ToSlice returns all items in the stack as a slice
func (s *Stack[T]) ToSlice() []T {
	return s.items
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}
