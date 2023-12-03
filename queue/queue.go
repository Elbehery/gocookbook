package queue

import "errors"

var (
	ErrEmptyQueue = errors.New("queue is empty")
	ErrNilQueue   = errors.New("queue is nil")
)

type Queue struct {
	elems []any
}

func NewQueue() *Queue {
	return &Queue{}
}

func newQueue(elems []any) *Queue {
	return &Queue{elems: elems}
}

func (q *Queue) Enqueue(elem any) error {
	if q == nil {
		return ErrNilQueue
	}
	q.elems = append(q.elems, elem)
	return nil
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}
	elem := q.elems[0]
	q.elems = q.elems[1:]
	return elem, nil
}

func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}
	return q.elems[0], nil
}

func (q *Queue) IsEmpty() bool {

	return q == nil || len(q.elems) == 0
}

func (q *Queue) Size() int {
	if q == nil {
		return 0
	}
	return len(q.elems)
}
