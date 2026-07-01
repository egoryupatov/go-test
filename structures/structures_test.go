package structures_test

import (
	"go_practice/structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Linked List

func TestLinkedListPushFront(t *testing.T) {
	l := &structures.LinkedList[int]{}

	l.PushFront(10)
	l.PushFront(20)

	assert.Equal(t, []int{20, 10}, l.All())
}

func TestLinkedListPushBack(t *testing.T) {
	l := &structures.LinkedList[int]{}

	l.PushBack(10)
	l.PushBack(20)

	assert.Equal(t, []int{10, 20}, l.All())
}

func TestLinkedListPopFront(t *testing.T) {
	l := &structures.LinkedList[int]{}

	l.PushBack(10)
	l.PushBack(20)

	v, ok := l.PopFront()

	assert.True(t, ok)
	assert.Equal(t, 10, v)
	assert.Equal(t, []int{20}, l.All())
}

func TestLinkedListPopBack(t *testing.T) {
	l := &structures.LinkedList[int]{}

	l.PushBack(10)
	l.PushBack(20)

	v, ok := l.PopBack()

	assert.True(t, ok)
	assert.Equal(t, 20, v)
	assert.Equal(t, []int{10}, l.All())
}

func TestLinkedListInsertAfter(t *testing.T) {
	l := &structures.LinkedList[int]{}

	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	l.InsertAfter(20, 25)

	assert.Equal(t, []int{10, 20, 25, 30}, l.All())
}

func TestLinkedListInsertBefore(t *testing.T) {
	l := &structures.LinkedList[int]{}

	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	l.InsertBefore(20, 15)

	assert.Equal(t, []int{10, 15, 20, 30}, l.All())
}

func TestLinkedListSearch(t *testing.T) {
	l := &structures.LinkedList[int]{}

	l.PushBack(10)
	l.PushBack(20)

	v, ok := l.Search(20)

	assert.True(t, ok)
	assert.Equal(t, 20, v)

	_, ok = l.Search(100)
	assert.False(t, ok)
}

func TestLinkedListRemoveAtIndex(t *testing.T) {
	l := &structures.LinkedList[int]{}

	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	ok := l.RemoveAtIndex(1)

	assert.True(t, ok)
	assert.Equal(t, []int{10, 30}, l.All())
}

func TestLinkedListClear(t *testing.T) {
	l := &structures.LinkedList[int]{}

	l.PushBack(10)
	l.PushBack(20)

	l.Clear()

	assert.Equal(t, []int{}, l.All())
}

func TestLinkedListRange(t *testing.T) {
	l := &structures.LinkedList[int]{}

	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	var result []int

	for v := range l.Range() {
		result = append(result, v)
	}

	assert.Equal(t, []int{10, 20, 30}, result)
}

// Stack

func TestStackPush(t *testing.T) {
	s := &structures.Stack[int]{}

	s.Push(10)
	s.Push(20)

	assert.Equal(t, 2, s.Size())
	assert.Equal(t, 20, s.Pop())
	assert.Equal(t, 10, s.Pop())
}

func TestStackPop(t *testing.T) {
	s := &structures.Stack[int]{}

	s.Push(50)
	s.Push(100)
	s.Push(150)

	assert.Equal(t, 150, s.Pop())
}

func TestStackPeek(t *testing.T) {
	s := &structures.Stack[int]{}

	s.Push(10)
	s.Push(20)

	assert.Equal(t, 20, s.Peek())
}

func TestStackSize(t *testing.T) {
	s := structures.Stack[int]{}

	for v := range 5 {
		s.Push(v)
	}

	assert.Equal(t, 5, s.Size())
}

func TestStackClear(t *testing.T) {
	s := structures.Stack[int]{}
	
	for v := range 5 {
		s.Push(v)
	}

	s.Clear()

	assert.Equal(t, 0, s.Size())
}

// Queue

func TestQueueEnqueue(t *testing.T) {
	q := structures.Queue[int]{}

	q.Enqueue(10)
	q.Enqueue(20)

	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 10, q.Peek())
}

func TestQueueDequeue(t *testing.T) {
	q := structures.Queue[int]{}

	for v := range 5 {
		q.Enqueue(v)
	}

	assert.Equal(t, 0, q.Dequeue())
	assert.Equal(t, 4, q.Size())
}

func TestQueuePeek(t *testing.T) {
	q := structures.Queue[int]{}

	for v := range 3 {
		q.Enqueue(v)
	}

	assert.Equal(t, 0, q.Peek())
}

func TestQueueSize(t *testing.T) {
	q := structures.Queue[int]{}

	for v := range 10 {
		q.Enqueue(v)
	}
	
	assert.Equal(t, 10, q.Size())
}

func TestQueueClear(t *testing.T) {
	q := structures.Queue[int]{}

	for v := range 10 {
		q.Enqueue(v)
	}

	q.Clear()

	assert.Equal(t, 0, q.Size())
}