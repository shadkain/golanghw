package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	sortlib "sort"
	"strings"
)

type FlagSet struct {
	f, u, r, n bool
	o          string
	k          int
}

func main() {
	defer func() {
		if perr := recover(); perr != nil {
			fmt.Printf("Error: %s\n", perr)
			os.Exit(1)
		}
	}()

	flags, filename := parseCommandLine()

	in := readFile(filename)
	out, err := sort(in, flags)
	if err != nil {
		panic(err)
	}

	printResult(out, flags.o)
}

func sort(in string, flags *FlagSet) (out string, err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = fmt.Errorf("%s", perr)
		}
	}()

	lines := strings.Split(in, "\n")
	cmp := buildComparator(flags)
	doSort(lines, cmp)
	if flags.u == true {
		lines = unique(lines, cmp)
	}
	out = uniteStrings(lines)

	return
}

func parseCommandLine() (flags *FlagSet, filename string) {
	flags = &FlagSet{}

	flag.BoolVar(&flags.f, "f", false, "ignore letter case")
	flag.BoolVar(&flags.u, "u", false, "print only the first among several equal")
	flag.BoolVar(&flags.r, "r", false, "decrease order sort")
	flag.StringVar(&flags.o, "o", "", "output to a file, without this option output to stdout")
	flag.BoolVar(&flags.n, "n", false, "sort numbers")
	flag.IntVar(&flags.k, "k", -1, "sort by column (column separator is a space)")

	flag.Parse()

	filename = getFilename()

	return
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

func buildComparator(flags *FlagSet) *Comparator {
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
	sortlib.SliceStable(s, func(i, j int) bool {
		return cmp.Less(s[i], s[j])
	})
}

func unique(s []string, cmp *Comparator) []string {
	sLen := len(s)

	newS := make([]string, 0, sLen)
	newS = append(newS, s[0])

	for i := 1; i < sLen; i++ {
		if !cmp.Equal(s[i], s[i-1]) {
			newS = append(newS, s[i])
		}
	}

	return newS
}

func uniteStrings(s []string) string {
	var res string

	for i := range s {
		res += s[i] + "\n"
	}

	res = res[:len(res)-1]

	return res
}

func printResult(str string, outfile string) {
	var stream *os.File
	var lastSym = ""
	if outfile == "" {
		stream = os.Stdin
		lastSym = "\n"
	} else {
		var err error
		stream, err = os.Create(outfile)
		if err != nil {
			panic(err)
		}

		defer stream.Close()
	}

	_, err := stream.WriteString(str + lastSym)
	if err != nil {
		panic(err)
	}
}
