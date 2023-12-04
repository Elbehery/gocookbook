package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap_Push(t *testing.T) {
	testCases := []struct {
		name string
		h    *MinHeap[int]
		e    int
		exp  *MinHeap[int]
	}{
		{
			name: "happy path",
			h:    newMinHeap([]int{1, 2, 5}),
			e:    3,
			exp:  newMinHeap([]int{1, 2, 5, 3}),
		},
		{
			name: "unHappy path",
			h:    newMinHeap([]int{10, 3, 5, 6}),
			e:    4,
			exp:  newMinHeap([]int{3, 4, 5, 10, 6}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.h.Push(tc.e)
			assert.Equal(t, tc.exp, tc.h)
		})
	}
}
