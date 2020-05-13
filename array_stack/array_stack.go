package array_stack

type ArrayStack struct {
	n, cap int
	buf    []int
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
	as.buf = append(as.buf[:i+1], as.buf[i:]...)
	as.buf[i] = v
}

func (as *ArrayStack) Remove(i int) int {
	r := as.buf[i]
	as.buf = append(as.buf[:i], as.buf[i+1:]...)
	if as.isSeparate() {
		as.resize()
	}
	return r
}

func (as ArrayStack) isSeparate() bool {
	return cap(as.buf) >= len(as.buf)*3
}

func (as ArrayStack) isFull() bool {
	return len(as.buf) == cap(as.buf)
}

func (as ArrayStack) resize() {
	c := as.maximum(len(as.buf)*2, 1)
	newBuf := make([]int, len(as.buf), c)
	for i := 0; i < len(as.buf); i++ {
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
