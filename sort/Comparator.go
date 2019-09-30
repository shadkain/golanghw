package main

import (
	"strconv"
	"strings"
)

/** Comparator */
type Comparator struct {
	head    Node
	tail    Node
	reverse bool
}

func (this *Comparator) ApplyNode(node Node) {
	this.tail.SetNext(node)
	this.tail = node
}

func (this *Comparator) prepare(l, r string) Node {
	var node = this.head
	node.PutIn(l, r)

	for ; node.Next() != nil; node = node.Next() {
		var next = node.Next()
		next.PutIn(node.PutOut())
	}

	return node
}

func (this *Comparator) Less(l, r string) bool {
	var node = this.prepare(l, r)

	if this.reverse == true {
		return !node.Less()
	}

	return node.Less()
}

func (this *Comparator) Equal(l, r string) bool {
	var node = this.prepare(l, r)

	return node.Equal()
}

/** Nodes */
type Node interface {
	PutIn(string, string)
	PutOut() (string, string)
	Less() bool
	Equal() bool
	Next() Node
	SetNext(Node)
}

type BaseNode struct {
	l, r     string
	nextNode Node
}

func (this *BaseNode) PutIn(l, r string) {
	this.l = l
	this.r = r
}

func (this *BaseNode) PutOut() (string, string) {
	return this.l, this.r
}

func (this *BaseNode) Less() bool {
	return this.l < this.r
}

func (this *BaseNode) Equal() bool {
	return this.l == this.r
}

func (this *BaseNode) Next() Node {
	return this.nextNode
}

func (this *BaseNode) SetNext(node Node) {
	this.nextNode = node
}

type ColumnNode struct {
	BaseNode
	k   int
	sep string
}

func (this *ColumnNode) PutIn(l, r string) {
	this.l = strings.Split(l, this.sep)[this.k]
	this.r = strings.Split(r, this.sep)[this.k]
}

type IgnoreCaseNode struct {
	BaseNode
}

func (this *IgnoreCaseNode) PutIn(l, r string) {
	this.l = strings.ToUpper(l)
	this.r = strings.ToUpper(r)
}

type NumberNode struct {
	BaseNode
	li, ri int
}

func (this *NumberNode) PutIn(l, r string) {
	this.BaseNode.PutIn(l, r)

	li, err := strconv.Atoi(this.l)
	if err != nil {
		panic(err)
	}

	ri, err := strconv.Atoi(this.r)
	if err != nil {
		panic(err)
	}

	this.li = li
	this.ri = ri
}

func (this *NumberNode) Less() bool {
	return this.li < this.ri
}

func (this *NumberNode) Equal() bool {
	return this.li == this.ri
}
