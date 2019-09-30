package main

type StringStack struct {
	buffer []string
}

func NewStringStack(size int) *StringStack {
	return &StringStack{
		buffer: make([]string, 0, size),
	}
}

func (this *StringStack) IsEmpty() bool {
	return len(this.buffer) == 0
}

func (this *StringStack) Push(value string) {
	this.buffer = append(this.buffer, value)
}

func (this *StringStack) Top() string {
	return this.buffer[len(this.buffer)-1]
}

func (this *StringStack) Pop() bool {
	size := len(this.buffer)

	if size > 0 {
		this.buffer = this.buffer[:size-1]
		return true
	}

	return false
}
