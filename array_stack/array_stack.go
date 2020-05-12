package array_stack

type ArrayStack struct {
	n, cap int
	buf []int
}

func New() ArrayStack {
	return ArrayStack{}
}

func (as ArrayStack) Size() int {
	return as.n
}

func (as ArrayStack) Get(i int) int {
	return as.buf[i]
}

func (as *ArrayStack) Set(i, v int) int {
	as.buf[i] = v
	return as.buf[i]
}

func (as *ArrayStack) Add(i, v int) {
	if as.isFull() {
		as.resize()
	}
	for s := as.n; s > i; s-- {
		as.buf[s] = as.buf[s-1]
	}
	as.buf[i] = v
	as.n++
}

func (as *ArrayStack) Remove(i int) int {
	r := as.buf[i]
	for s := i; s < as.n - 1; s++ {
		as.buf[s] = as.buf[s + 1]
	}
	as.n--
	if as.isSeparate() {

		as.resize()
	}
	return r
}

func (as ArrayStack) isSeparate() bool {
	return len(as.buf) >= as.n * 3
}

func (as ArrayStack) isFull() bool {
	return as.n == as.cap
}

func (as ArrayStack) resize() {
	as.cap = as.maximum(as.n * 2, 1)
	newBuf := make([]int, as.cap)
	for i := 0; i < as.n; i++ {
		newBuf[i] = as.buf[i]
	}
	as.buf = newBuf
}

func (as ArrayStack) maximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}