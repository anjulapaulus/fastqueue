package fastqueue

import (
	"fmt"
	"testing"
)

func TestQueue_Push(t *testing.T) {
	var q Queue
	q.Push(0, 4)
	q.Push(1, 2)
	q.Push(2, 2)
	q.Push(3, 1)
	if q.Len() != 4 {
		t.Error("Push Function Failed")
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
func TestQueue_NullPeek(t *testing.T) {
	var q Queue

	if q.Peek() != 0 {
		t.Error("Peek Function Null check Failed")
	}
}
func TestQueue_PeekValue(t *testing.T) {
	var q Queue
	q.Push(0, 1)
	q.Push(1, 2)

	if q.PeekValue() != float64(1) {
		t.Error("PeekValue function Null check failed")
	}
}
func TestQueue_NullPeekValue(t *testing.T) {
	var q Queue

	if q.PeekValue() != float64(0) {
		t.Error("PeekValue function failed")
	}
}

func TestQueue_Pop(t *testing.T) {
	var q Queue
	q.Push(0, 1)
	q.Push(1, 2)
	q.Push(2, 2)
	q.Push(3, 6)
	q.Push(4, 5)
	q.Push(5, 1)
	q.Pop()
	q.Pop()
	id, val := q.Pop()
	if id != 2 && val != float64(2) {
		t.Error("Pop function failed")
	}
}
func TestQueue_NullPop(t *testing.T) {
	var q Queue

	id, val := q.Pop()
	if id != 0 && val != float64(0) {
		t.Error("Pop function Null Check failed")
	}
}
func TestQueue_All(t *testing.T) {
	var q Queue
	q.Push(0, 1)
	q.Push(1, 2)

	mapp := q.All()
	fmt.Println(mapp)

	if mapp == nil {
		t.Error("All function failed")
	}
}
