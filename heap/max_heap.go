package heap

import "golang.org/x/exp/constraints"

type MaxHeap[T constraints.Ordered] struct {
	nodes []T
}

func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{}
}

func newMaxHeap[T constraints.Ordered](nodes []T) *MaxHeap[T] {
	h := NewMaxHeap[T]()
	for _, v := range nodes {
		h.Push(v)
	}
	return h
}

func (h *MaxHeap[T]) Push(val T) {
	h.nodes = append(h.nodes, val)
	cur := len(h.nodes) - 1
	for ; h.nodes[cur] > h.nodes[parent(cur)]; cur = parent(cur) {
		h.Swap(cur, parent(cur))
	}
}

func (h *MaxHeap[T]) Pop() T {
	res := h.nodes[0]
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.reArrange(0)
	return res
}

func (h *MaxHeap[T]) reArrange(i int) {
	largest := i
	left, right := leftChild(i), rightChild(i)

	if left < h.Size() && h.nodes[left] > h.nodes[largest] {
		largest = left
	}

	if right < h.Size() && h.nodes[right] > h.nodes[largest] {
		largest = right
	}

	if largest != i {
		h.Swap(i, largest)
		h.reArrange(largest)
	}
}

func (h *MaxHeap[T]) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *MaxHeap[T]) IsEmpty() bool {
	return len(h.nodes) == 0
}

func (h *MaxHeap[T]) Size() int {
	return len(h.nodes)
}
