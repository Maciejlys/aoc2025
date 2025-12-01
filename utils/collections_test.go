package utils

import "testing"

func TestSet(t *testing.T) {
	s := NewSet[int]()
	
	if s.Size() != 0 {
		t.Errorf("New set size = %d, want 0", s.Size())
	}
	
	s.Add(1)
	s.Add(2)
	s.Add(3)
	
	if s.Size() != 3 {
		t.Errorf("Set size = %d, want 3", s.Size())
	}
	
	if !s.Contains(2) {
		t.Error("Set should contain 2")
	}
	
	s.Remove(2)
	
	if s.Contains(2) {
		t.Error("Set should not contain 2 after removal")
	}
	
	if s.Size() != 2 {
		t.Errorf("Set size = %d, want 2", s.Size())
	}
}

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
	
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	
	if q.Size() != 3 {
		t.Errorf("Queue size = %d, want 3", q.Size())
	}
	
	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("Dequeue() = %d, %v, want 1, true", val, ok)
	}
	
	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Errorf("Dequeue() = %d, %v, want 2, true", val, ok)
	}
	
	if q.Size() != 1 {
		t.Errorf("Queue size = %d, want 1", q.Size())
	}
}

func TestStack(t *testing.T) {
	s := NewStack[int]()
	
	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}
	
	s.Push(1)
	s.Push(2)
	s.Push(3)
	
	if s.Size() != 3 {
		t.Errorf("Stack size = %d, want 3", s.Size())
	}
	
	val, ok := s.Pop()
	if !ok || val != 3 {
		t.Errorf("Pop() = %d, %v, want 3, true", val, ok)
	}
	
	val, ok = s.Pop()
	if !ok || val != 2 {
		t.Errorf("Pop() = %d, %v, want 2, true", val, ok)
	}
	
	if s.Size() != 1 {
		t.Errorf("Stack size = %d, want 1", s.Size())
	}
}
