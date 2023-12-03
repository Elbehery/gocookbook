package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_IsEmpty(t *testing.T) {
	testCases := []struct {
		name string
		s    *Stack
		exp  bool
	}{
		{
			name: "length 5",
			s:    newStack([]any{1, 2, 3, 4, 5}),
			exp:  false,
		},
		{
			name: "length 1",
			s:    newStack([]any{"x"}),
			exp:  false,
		},
		{
			name: "empty",
			s:    newStack([]any{}),
			exp:  true,
		},
		{
			name: "nil stack",
			s:    nil,
			exp:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.s.IsEmpty()
			assert.Equal(t, tc.exp, act)
		})
	}
}

func TestStack_Push(t *testing.T) {
	testCases := []struct {
		name   string
		s      *Stack
		elem   any
		expErr error
	}{
		{
			name:   "empty stack",
			s:      newStack([]any{}),
			elem:   1,
			expErr: nil,
		},
		{
			name:   "nil stack",
			s:      nil,
			elem:   "x",
			expErr: ErrNilStack,
		},
		{
			name:   "stack of ints",
			s:      newStack([]any{1, 2, 3, 4}),
			elem:   5,
			expErr: nil,
		},
		{
			name:   "stack of strings",
			s:      newStack([]any{"x", "y"}),
			elem:   "z",
			expErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.s.Push(tc.elem)
			assert.Equal(t, tc.expErr, act)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	testCases := []struct {
		name   string
		s      *Stack
		exp    any
		expErr error
	}{
		{
			name:   "nil stack",
			s:      nil,
			exp:    nil,
			expErr: ErrEmptyStack,
		},
		{
			name:   "empty stack",
			s:      newStack([]any{}),
			exp:    nil,
			expErr: ErrEmptyStack,
		},
		{
			name:   "stack of ints",
			s:      newStack([]any{1, 2, 3, 4, 5}),
			exp:    5,
			expErr: nil,
		},
		{
			name:   "stack of a string",
			s:      newStack([]any{"x"}),
			exp:    "x",
			expErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act, err := tc.s.Pop()
			assert.Equal(t, tc.expErr, err)
			assert.Equal(t, tc.exp, act)
		})
	}
}

func TestStack_Peek(t *testing.T) {
	testCases := []struct {
		name    string
		s       *Stack
		exp     any
		expErr  error
		expSize int
	}{
		{
			name:   "nil stack",
			s:      nil,
			exp:    nil,
			expErr: ErrEmptyStack,
		},
		{
			name:   "empty Stack",
			s:      newStack([]any{}),
			exp:    nil,
			expErr: ErrEmptyStack,
		},
		{
			name:   "stack of ints",
			s:      newStack([]any{1, 2, 3, 4, 5}),
			exp:    5,
			expErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act, err := tc.s.Peek()
			assert.Equal(t, tc.expErr, err)
			assert.Equal(t, tc.exp, act)
		})
	}
}
