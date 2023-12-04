package heap

import "golang.org/x/exp/constraints"

type MinHeap[T constraints.Ordered] struct {
	nodes []T
}

func NewMinHeap[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{}
}

func newMinHeap[T constraints.Ordered](nodes []T) *MinHeap[T] {
	return &MinHeap[T]{nodes: nodes}
}

func (h *MinHeap[T]) Push(val T) {
	h.nodes = append(h.nodes, val)
	for cur := len(h.nodes) - 1; cur < parent(cur); cur = parent(cur) {
		h.swap(cur, parent(cur))
	}
}

func (h *MinHeap[T]) Pop() T {
	res := h.nodes[0]
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.reArrange(0)
	return res
}

func (h *MinHeap[T]) reArrange(i int) {
	minimal := i
	left, right := leftChild(i), rightChild(i)

	if left < h.Size() && h.nodes[left] < h.nodes[minimal] {
		minimal = left
	}

	if right < h.Size() && h.nodes[right] < h.nodes[minimal] {
		minimal = right
	}

	if minimal != i {
		h.swap(i, minimal)
		h.reArrange(minimal)
	}
}

func (h *MinHeap[T]) Size() int {
	return len(h.nodes)
}

func (h *MinHeap[T]) IsEmpty() bool {
	return len(h.nodes) == 0
}

func (h *MinHeap[T]) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}
func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return (2 * i) + 1
}

func rightChild(i int) int {
	return (2 * i) + 2
}
