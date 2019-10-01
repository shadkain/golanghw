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
	ignoreCase, unique, reverse, numbers bool
	output                               string
	column                               int
}

func main() {
	flags := &FlagSet{}
	filename := ""

	err := parseCommandLine(flags, &filename)
	if err != nil {
		crash(err)
	}

	in := strings.Split(readFile(filename), "\n")
	out, err := sort(in, flags)
	if err != nil {
		crash(err)
	}

	err = printResult(out, flags.output)
	if err != nil {
		crash(err)
	}
}

func crash(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}

func sort(in []string, flags *FlagSet) (out []string, err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = fmt.Errorf("%s", perr)
		}
	}()

	lines := in
	cmp := buildComparator(flags)
	doSort(lines, cmp)
	if flags.unique == true {
		lines = unique(lines, cmp)
	}
	out = lines

	return
}

func parseCommandLine(flags *FlagSet, filename *string) (err error) {
	flags = &FlagSet{}

	flag.BoolVar(&flags.ignoreCase, "f", false, "ignore letter case")
	flag.BoolVar(&flags.unique, "u", false, "print only the first among several equal")
	flag.BoolVar(&flags.reverse, "r", false, "decrease order sort")
	flag.StringVar(&flags.output, "o", "", "output to a file, without this option output to stdout")
	flag.BoolVar(&flags.numbers, "n", false, "sort numbers")
	flag.IntVar(&flags.column, "k", -1, "sort by column (column separator is a space)")

	flag.Parse()

	*filename, err = getFilename()

	return
}

func getFilename() (string, error) {
	if flag.NArg() < 1 {
		return "", fmt.Errorf("input file is missing")
	}

	return flag.Arg(0), nil
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
		reverse: flags.reverse,
		head:    base,
		tail:    base,
	}

	if flags.column > -1 {
		cmp.ApplyNode(&ColumnNode{
			k:   flags.column,
			sep: " ",
		})
	}

	if flags.numbers == true {
		cmp.ApplyNode(&NumberNode{})
	} else if flags.ignoreCase == true {
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

func printResult(lines []string, outfile string) error {
	var stream *os.File
	if outfile == "" {
		stream = os.Stdin
	} else {
		var err error
		stream, err = os.Create(outfile)
		if err != nil {
			return err
		}

		defer stream.Close()
	}

	_, err := stream.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		return err
	}

	return nil
}
