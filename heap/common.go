package heap

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return (i * 2) + 1
}

func rightChild(i int) int {
	return (i * 2) + 2
}
