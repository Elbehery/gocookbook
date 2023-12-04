package linkedlist

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type Element[T constraints.Ordered] struct {
	val  T
	next *Element[T]
}

type LinkedList[T constraints.Ordered] struct {
	head *Element[T]
	size int
}

func NewLinkedList[T constraints.Ordered]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		size: 0,
	}
}

func (l *LinkedList[T]) Add(e *Element[T]) {
	if l.head == nil {
		l.head = e
	} else {
		e.next = l.head
		l.head = e
	}
	l.size++
}

func (l *LinkedList[T]) Insert(e *Element[T], marker T) error {
	for cur := l.head; cur.next != nil; cur = cur.next {
		if cur.val == marker {
			e.next = cur.next
			cur.next = e
			l.size++
			return nil
		}
	}
	return errors.New("elem not found")
}

func (l *LinkedList[T]) Delete(e *Element[T]) error {
	prev := l.head
	cur := l.head

	for cur != nil {
		if cur.val == e.val {
			if cur == l.head {
				l.head = cur.next
			} else {
				prev.next = cur.next
			}
			l.size--
			return nil
		}
		prev = cur
		cur = cur.next
	}
	return errors.New("elem not found")
}

func (l *LinkedList[T]) Find(val T) (*Element[T], error) {
	var res *Element[T]

	for cur := l.head; cur.next != nil; cur = cur.next {
		if cur.val == val {
			res = cur
		}
	}

	if res != nil {
		return res, nil
	}

	return nil, errors.New("elem not found")
}

func (l *LinkedList[T]) List() []*Element[T] {
	var res []*Element[T]

	for cur := l.head; cur != nil; cur = cur.next {
		res = append(res, cur)
	}
	return res
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Size() int {
	return l.size
}
