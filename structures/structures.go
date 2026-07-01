package structures

import "iter"

type Node[T comparable] struct {
	value T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type Stack[T any] struct {
	data []T
}

type Queue[T any] struct {
	data []T
}

// Linked List

func (l *LinkedList[T]) PushFront(elem T) {
	newNode := &Node[T]{
		value: elem,
		next: l.head,
	}

	l.head = newNode

	if l.tail == nil {
		l.tail = newNode
	}

	l.size++
}

func (l *LinkedList[T]) PushBack(elem T) {
	newNode := &Node[T]{
		value: elem,
		next: nil,
	}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return
	}

	l.tail.next = newNode
	l.tail = newNode
	l.size++
}

func (l *LinkedList[T]) PopFront() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}

	value := l.head.value
	l.head = l.head.next
	l.size--

	return value, true
}

func (l *LinkedList[T]) PopBack() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}

	if l.head.next == nil {
		value := l.head.value
		l.head = nil
		l.tail = nil
		l.size--
		return value, true
	}

	prev := l.head

	for prev.next.next != nil {
		prev = prev.next
	}

	value := prev.next.value
	prev.next = nil
	l.tail = prev

	l.size--

	return value, true
}

func (l *LinkedList[T]) InsertAfter(afterValue T, newValue T) {

	current := l.head

	for current != nil && current.value != afterValue {
		current = current.next
	}

	if current == nil {
		return
	}

	newNode := &Node[T]{
		value: newValue,
		next: current.next,
	}

	current.next = newNode

	if current == l.tail {
		l.tail = newNode
	}

	l.size++
} 

func (l *LinkedList[T]) InsertBefore(beforeValue T, newValue T) {

	if l.head == nil {
		return
	}

	if l.head.value == beforeValue {
		l.PushFront(newValue)
		return
	}

	current := l.head

	for current != nil && current.next.value != beforeValue {
		current = current.next
	}

	if current == nil {
		return
	}

	newNode := &Node[T]{
		value: newValue,
		next: current.next,
	}

	current.next = newNode
	l.size++
}

func (l *LinkedList[T]) Search(elem T) (T, bool) {

	current := l.head

	for current != nil {
		if (current.value == elem){
			return current.value, true
		}

		current = current.next
	}

	var zero T
	
	return zero, false
}

func (l *LinkedList[T]) RemoveAtIndex(index int) bool {
	if index < 0 || l.head == nil {
		return false
	}

	if index == 0 {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		}
		l.size--
		return true
	}

	current := l.head

	for i := 0; i < index-1; i++ {
		if current == nil {
			return false
		}
		current = current.next
	}

	if current == nil || current.next == nil {
		return false
	}

	current.next = current.next.next

	if current.next == nil {
		l.tail = current
	}

	l.size--

	return true
}

func (l *LinkedList[T]) Clear(){
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList[T]) All() []T {

	all := make([]T, 0, l.size)

	current := l.head

	for current != nil {
		all = append(all, current.value)
		
		current = current.next
	}

	return  all
}

func (l *LinkedList[T]) Range() iter.Seq[T] {
	return func(yield func(T) bool) {
		current := l.head

		for current != nil {
			if !yield(current.value) {
				return
			}
			current = current.next
		}
	}
}

// Stack

func (s *Stack[T]) Push(elem T) {
	s.data = append(s.data, elem)
}

func (s *Stack[T]) Pop() T {
	last := s.data[len(s.data) - 1]

	s.data = s.data[:len(s.data)-1]

	return last
}

func (s *Stack[T]) Peek() T {
	last := s.data[len(s.data) - 1]

	return last
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Clear() {
	s.data = make([]T, 0)
}

// Queue

func (q *Queue[T]) Enqueue(elem T) {
	q.data = append(q.data, elem)
}

func (q *Queue[T]) Dequeue() T {
	elem := q.data[0]
	q.data = q.data[1:]

	return elem
}

func (q *Queue[T]) Peek() T {
	elem := q.data[0]

	return elem
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) Clear() {
	q.data = make([]T, 0)
}


