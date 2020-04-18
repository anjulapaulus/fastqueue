package fastqueue

import (
	"testing"
)

func TestQueue_Push(t *testing.T) {
	var q Queue
	q.Push(0, 1)
	q.Push(1, 2)
	if q.Len() != 2 {
		t.Error("Push Function Failed")
	}
}

func BenchmarkQueue_Push(b *testing.B) {
	var q Queue
	for i := 0; i < b.N; i++ {
		q.Push(0, float64(i))
	}
}

func TestQueue_Len(t *testing.T) {
	var q Queue
	q.Push(0, 1)
	q.Push(1, 2)
	if q.Len() != 2 {
		t.Error("Push Function Failed")
	}
}

func TestQueue_Peek(t *testing.T) {
	var q Queue
	q.Push(0, 1)
	q.Push(1, 2)

	if q.Peek() != 0 {
		t.Error("Peek Function Failed")
	}
}

func TestQueue_PeekValue(t *testing.T) {
	var q Queue
	q.Push(0, 1)
	q.Push(1, 2)

	if q.PeekValue() != float64(1) {
		t.Error("Pop function failed")
	}
}

func TestQueue_Pop(t *testing.T) {
	var q Queue
	q.Push(0, 1)
	q.Push(1, 2)

	id, val := q.Pop()
	if id != 0 && val != float64(1) {
		t.Error("Pop function failed")
	}
}
