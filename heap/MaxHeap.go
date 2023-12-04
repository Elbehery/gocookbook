package heap

import "golang.org/x/exp/constraints"

type MaxHeap[T constraints.Ordered] struct {
	nodes []T
}

func (h *MaxHeap[T]) NewMaxHeap() *MaxHeap[T] {
	return &MaxHeap[T]{}
}

func (h *MaxHeap[T]) newMaxHeap(nodes []T) *MaxHeap[T] {
	return &MaxHeap[T]{nodes: nodes}
}

func (h *MaxHeap[T]) Push(val T) {
	h.nodes = append(h.nodes, val)
	for cur := len(h.nodes) - 1; h.nodes[cur] > h.nodes[h.parent(cur)]; cur = h.parent(cur) {
		h.swap(cur, h.parent(cur))
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
	left, right := h.leftChild(i), h.rightChild(i)

	if left < h.Size() && h.nodes[left] > h.nodes[largest] {
		largest = left
	}

	if right < h.Size() && h.nodes[right] > h.nodes[largest] {
		largest = right
	}

	if largest != i {
		h.swap(i, largest)
		h.reArrange(largest)
	}
}

func (h *MaxHeap[T]) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *MaxHeap[T]) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap[T]) leftChild(i int) int {
	return (i * 2) + 1
}

func (h *MaxHeap[T]) rightChild(i int) int {
	return (i * 2) + 2
}

func (h *MaxHeap[T]) Size() int {
	return len(h.nodes)
}

func (h *MaxHeap[T]) IsEmpty() bool {
	return len(h.nodes) == 0
}
