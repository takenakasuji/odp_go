package array_queue

type ArrayQueue struct {
	n, cap, j int
	buf       []int
}

func New() ArrayQueue {
	return ArrayQueue{}
}

func (as *ArrayQueue) Add(v int) {
	if as.isFull() {
		as.resize()
	}
	as.buf[(as.j+as.n)%as.cap] = v
	as.n++
}

func (as *ArrayQueue) Remove() int {
	ret := as.buf[as.j]
	as.j = (as.j + 1) % as.cap
	as.n--
	if as.isSparse() {
		as.resize()
	}
	return ret
}

func (as ArrayQueue) isFull() bool {
	return as.n == as.cap
}

func (as ArrayQueue) isSparse() bool {
	return len(as.buf) >= 3*as.n
}

func (as *ArrayQueue) resize() {
	capNew := as.maximum(2*as.n, 1)
	bufNew := make([]int, capNew)
	for i := 0; i < as.n; i++ {
		bufNew[i] = as.buf[(i+as.j)%as.cap]
	}
	as.buf = bufNew
	as.cap = capNew
	as.j = 0
}

func (as ArrayQueue) maximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
