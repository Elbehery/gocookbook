package stack

import "errors"

var (
	ErrEmptyStack = errors.New("stack is empty")
	ErrNilStack   = errors.New("stack is nil")
)

type Stack struct {
	elems []any
}

func NewStack() *Stack {
	return &Stack{}
}

func newStack(elems []any) *Stack {
	return &Stack{elems: elems}
}

func (s *Stack) Push(elem any) error {
	if s == nil {
		return ErrNilStack
	}
	s.elems = append(s.elems, elem)
	return nil
}

func (s *Stack) IsEmpty() bool {
	return s == nil || len(s.elems) == 0
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem, nil
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}
	e := s.elems[len(s.elems)-1]
	return e, nil
}

func (s *Stack) Size() int {
	return len(s.elems)
}
