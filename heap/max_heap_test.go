package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap_Push(t *testing.T) {
	testCases := []struct {
		name string
		h    *MaxHeap[int]
		e    int
		exp  *MaxHeap[int]
	}{
		{
			name: "happy path",
			h:    newMaxHeap([]int{1, 2, 5}),
			e:    3,
			exp:  newMaxHeap([]int{5, 3, 2, 1}),
		},
		{
			name: "unHappy path",
			h:    newMaxHeap([]int{4, 7, 3, 10}),
			e:    5,
			exp:  newMaxHeap([]int{10, 7, 3, 4, 5}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.h.Push(tc.e)
			assert.Equal(t, tc.exp, tc.h)
		})
	}
}

func TestMaxHeap_Pop(t *testing.T) {
	testCases := []struct {
		name string
		h    *MaxHeap[int]
		exp  int
	}{
		{
			name: "happy path",
			h:    newMaxHeap([]int{1, 2, 5}),
			exp:  5,
		},
		{
			name: "unHappy path",
			h:    newMaxHeap([]int{4, 7, 3, 10}),
			exp:  10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.h.Pop()
			assert.Equal(t, tc.exp, act)
		})
	}
}
