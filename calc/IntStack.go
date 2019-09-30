package main

type IntStack struct {
	buffer []int
}

func NewIntStack(size int) *IntStack {
	return &IntStack{
		buffer: make([]int, 0, size),
	}
}

func (this *IntStack) IsEmpty() bool {
	return len(this.buffer) == 0
}

func (this *IntStack) Push(value int) {
	this.buffer = append(this.buffer, value)
}

func (this *IntStack) Top() int {
	return this.buffer[len(this.buffer)-1]
}

func (this *IntStack) Pop() bool {
	size := len(this.buffer)

	if size > 0 {
		this.buffer = this.buffer[:size-1]
		return true
	}

	return false
}
