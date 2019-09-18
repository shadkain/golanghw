package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flagSet struct {
	f, u, r, n bool
	o          string
	k          int
}

func initFlags(flags *flagSet) {
	flag.BoolVar(&flags.f, "f", false, "ignore letter case")
	flag.BoolVar(&flags.u, "u", false, "print only the first among several equal")
	flag.BoolVar(&flags.r, "r", false, "decrease order sort")
	flag.StringVar(&flags.o, "o", "", "output to a file, without this option output to stdout")
	flag.BoolVar(&flags.n, "n", false, "sort numbers")
	flag.IntVar(&flags.k, "k", -1, "sort by column (column separator is a space)")
}

func getFilename() string {
	if flag.NArg() < 1 {
		panic("input file is missing")
	}

	return flag.Arg(0)
}

func readFile(filename string) string {
	var data, err = ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func buildComparator(flags *flagSet) *Comparator {
	var base = &BaseNode{}

	var cmp = &Comparator{
		reverse: flags.r,
		head:    base,
		tail:    base,
	}

	if flags.k > -1 {
		cmp.ApplyNode(&ColumnNode{
			k:   flags.k,
			sep: " ",
		})
	}

	if flags.n == true {
		cmp.ApplyNode(&NumberNode{})
	} else if flags.f == true {
		cmp.ApplyNode(&IgnoreCaseNode{})
	}

	return cmp
}

func doSort(s []string, cmp *Comparator) {
	sort.SliceStable(s, func(i, j int) bool {
		return cmp.Less(s[i], s[j])
	})
}

func printResult(s []string, cmp *Comparator, flags *flagSet) {
	var stream *os.File
	if flags.o == "" {
		stream = os.Stdin
	} else {
		var err error
		stream, err = os.Create(flags.o)
		if err != nil {
			panic(err)
		}

		defer stream.Close()
	}

	var sLen = len(s)
	var addSym string
	for i := range s {
		if i > 1 && flags.u == true {
			if cmp.Equal(s[i], s[i-1]) {
				continue
			}
		}

		if i+1 < sLen {
			addSym = "\n"
		} else {
			addSym = ""
		}

		_, err := stream.WriteString(s[i] + addSym)
		if err != nil {
			panic(err)
		}
	}

	if flags.u == true {

	}
}

func errorProcess() {
	var err = recover()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func main() {
	var flags = &flagSet{}
	initFlags(flags)
	flag.Parse()

	defer errorProcess()
	var s = strings.Split(readFile(getFilename()), "\n")
	var cmp = buildComparator(flags)
	doSort(s, cmp)
	printResult(s, cmp, flags)
}

// Comparator with nodes
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
