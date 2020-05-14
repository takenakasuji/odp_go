package array_queue

type ArrayQueue struct {
	buf []int
}

func New() ArrayQueue {
	return ArrayQueue{}
}
