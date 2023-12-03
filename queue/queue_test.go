package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_IsEmpty(t *testing.T) {
	testCases := []struct {
		name string
		q    *Queue
		exp  bool
	}{
		{
			name: "5 elems",
			q:    newQueue([]any{1, 2, 3, 4, 5}),
			exp:  false,
		},
		{
			name: "1 elem",
			q:    newQueue([]any{"x"}),
			exp:  false,
		},
		{
			name: "empty queue",
			q:    NewQueue(),
			exp:  true,
		},
		{
			name: "nil queue",
			q:    nil,
			exp:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.q.IsEmpty()
			assert.Equal(t, tc.exp, act)
		})
	}
}

func TestQueue_Size(t *testing.T) {
	testCases := []struct {
		name string
		q    *Queue
		exp  int
	}{
		{
			name: "5 elems",
			q:    newQueue([]any{1, 2, 3, 4, 5}),
			exp:  5,
		},
		{
			name: "1 elem",
			q:    newQueue([]any{"x"}),
			exp:  1,
		},
		{
			name: "empty queue",
			q:    NewQueue(),
			exp:  0,
		},
		{
			name: "nil queue",
			q:    nil,
			exp:  0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.q.Size()
			assert.Equal(t, tc.exp, act)
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	testCases := []struct {
		name   string
		q      *Queue
		elem   any
		exp    *Queue
		expErr error
	}{
		{
			name:   "queue of ints",
			q:      newQueue([]any{1, 2, 3, 4}),
			elem:   5,
			exp:    newQueue([]any{1, 2, 3, 4, 5}),
			expErr: nil,
		},
		{
			name:   "queue of strings",
			q:      newQueue([]any{"x", "y"}),
			elem:   "z",
			exp:    newQueue([]any{"x", "y", "z"}),
			expErr: nil,
		},
		{
			name:   "nil queue",
			q:      nil,
			elem:   1,
			exp:    nil,
			expErr: ErrNilQueue,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.q.Enqueue(tc.elem)
			assert.Equal(t, tc.expErr, err)
			assert.Equal(t, tc.exp, tc.q)
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	testCases := []struct {
		name   string
		q      *Queue
		expErr error
		exp    any
	}{
		{
			name:   "int elems",
			q:      newQueue([]any{1, 2, 3, 4, 5}),
			exp:    1,
			expErr: nil,
		},
		{
			name:   "string elem",
			q:      newQueue([]any{"x"}),
			exp:    "x",
			expErr: nil,
		},
		{
			name:   "empty queue",
			q:      NewQueue(),
			exp:    nil,
			expErr: ErrEmptyQueue,
		},
		{
			name:   "nil queue",
			q:      nil,
			exp:    nil,
			expErr: ErrEmptyQueue,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act, err := tc.q.Dequeue()
			assert.Equal(t, tc.expErr, err)
			assert.Equal(t, tc.exp, act)
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	testCases := []struct {
		name     string
		q        *Queue
		expErr   error
		exp      any
		expQueue *Queue
	}{
		{
			name:     "int elems",
			q:        newQueue([]any{1, 2, 3, 4, 5}),
			exp:      1,
			expErr:   nil,
			expQueue: newQueue([]any{1, 2, 3, 4, 5}),
		},
		{
			name:     "string elem",
			q:        newQueue([]any{"x"}),
			exp:      "x",
			expErr:   nil,
			expQueue: newQueue([]any{"x"}),
		},
		{
			name:     "empty queue",
			q:        NewQueue(),
			exp:      nil,
			expErr:   ErrEmptyQueue,
			expQueue: NewQueue(),
		},
		{
			name:     "nil queue",
			q:        nil,
			exp:      nil,
			expErr:   ErrEmptyQueue,
			expQueue: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act, err := tc.q.Peek()
			assert.Equal(t, tc.expErr, err)
			assert.Equal(t, tc.exp, act)
			assert.Equal(t, tc.expQueue, tc.q)
		})
	}
}
