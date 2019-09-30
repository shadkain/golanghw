package main

import (
	"testing"
)

/** Sort1 */
var s1in = `Soda
Apple
Join
Unite
Book
Debian
BOOK`

var s1ex = `Apple
BOOK
Book
Debian
Join
Soda
Unite`

func TestSort1(t *testing.T) {
	in := s1in
	expected := s1ex

	flags := &FlagSet{}

	testOk(in, expected, flags, t)
}

/** SortReverse1 */
var sr1in = `Soda
Apple
Join
Unite
Book
Debian
BOOK`

var sr1ex = `Unite
Soda
Join
Debian
Book
BOOK
Apple`

func TestSortReverse1(t *testing.T) {
	in := sr1in
	expected := sr1ex

	flags := &FlagSet{
		r: true,
	}

	testOk(in, expected, flags, t)
}

/** Sort2 */
var s2in = `7Up
1Helloworld
4From
1Looser
0Urgent`

var s2ex = `0Urgent
1Helloworld
1Looser
4From
7Up`

func TestSort2(t *testing.T) {
	in := s2in
	expected := s2ex

	flags := &FlagSet{}

	testOk(in, expected, flags, t)
}

/** IgnoreCase1 */
var ic1in = `Join
jOhn
ko13
KO09
Lolipop`

var ic1ex = `jOhn
Join
KO09
ko13
Lolipop`

func TestIgnoreCase1(t *testing.T) {
	in := ic1in
	expected := ic1ex

	flags := &FlagSet{
		f: true,
	}

	testOk(in, expected, flags, t)
}

/** IgnoreCaseReverse1 */
var icr1in = `Join
jOhn
ko13
KO09
Lolipop`

var icr1ex = `Lolipop
ko13
KO09
Join
jOhn`

func TestIgnoreCaseReverse1(t *testing.T) {
	in := icr1in
	expected := icr1ex

	flags := &FlagSet{
		f: true,
		r: true,
	}

	testOk(in, expected, flags, t)
}

/** IgnoreCaseUnique1 */
var icu1in = `JOHN
Join
jOhn
ko13
KO09
Lolipop`

var icu1ex = `JOHN
Join
KO09
ko13
Lolipop`

func TestIgnoreCaseUnique1(t *testing.T) {
	in := icu1in
	expected := icu1ex

	flags := &FlagSet{
		f: true,
		u: true,
	}

	testOk(in, expected, flags, t)
}

/** IgnoreCaseUniqueReverse1 */
var icur1in = `Join
jOhn
ko13
JOHN
KO09
Lolipop`

var icur1ex = `Lolipop
ko13
KO09
Join
JOHN`

func TestIgnoreCaseUniqueReverse1(t *testing.T) {
	in := icur1in
	expected := icur1ex

	flags := &FlagSet{
		f: true,
		u: true,
		r: true,
	}

	testOk(in, expected, flags, t)
}

/** Unique1 */
var u1in = `Lol
Bom
Kohn
Bom
Tideman`

var u1ex = `Bom
Kohn
Lol
Tideman`

func TestUnique1(t *testing.T) {
	in := u1in
	expected := u1ex

	flags := &FlagSet{
		u: true,
	}

	testOk(in, expected, flags, t)
}

/** UniqueReverse1 */
var ur1in = `Lol
Bom
Kohn
Bom
Tideman`

var ur1ex = `Tideman
Lol
Kohn
Bom`

func TestUniqueReverse1(t *testing.T) {
	in := ur1in
	expected := ur1ex

	flags := &FlagSet{
		u: true,
		r: true,
	}

	testOk(in, expected, flags, t)
}

/** Numbers1 */
var n1in = `9
76
12
67
994
-234`

var n1ex = `-234
9
12
67
76
994`

func TestNumbers1(t *testing.T) {
	in := n1in
	expected := n1ex

	flags := &FlagSet{
		n: true,
	}

	testOk(in, expected, flags, t)
}

/** NumbersReverse1 */
var nr1in = `9
76
12
67
994
-234`

var nr1ex = `994
76
67
12
9
-234`

func TestNumbersReverse1(t *testing.T) {
	in := nr1in
	expected := nr1ex

	flags := &FlagSet{
		n: true,
		r: true,
	}

	testOk(in, expected, flags, t)
}

/** Columns1 */
var c1in = `Jon Snow
Jorah Marmont
Deyneris Targarien
Aria Stark`

var c1ex = `Jorah Marmont
Jon Snow
Aria Stark
Deyneris Targarien`

func TestColumns1(t *testing.T) {
	in := c1in
	expected := c1ex

	flags := &FlagSet{
		k: 1,
	}

	testOk(in, expected, flags, t)
}

/** ColumnsNumbers1 */
var cn1in = `923 8 122
87 23 87
-98 23 9234844
221 764 -23
8765 -90 44
23123 -123 987`

var cn1ex = `221 764 -23
8765 -90 44
87 23 87
923 8 122
23123 -123 987
-98 23 9234844`

func TestColumnsNumbers1(t *testing.T) {
	in := cn1in
	expected := cn1ex

	flags := &FlagSet{
		k: 2,
		n: true,
	}

	testOk(in, expected, flags, t)
}

/** ColumnsIgnoreCase1 */
var cic1in = `1 Kok
1 KOj
1 kOa
1 koh
1 koB`

var cic1ex = `1 kOa
1 koB
1 koh
1 KOj
1 Kok`

func TestColumnsIgnoreCase1(t *testing.T) {
	in := cic1in
	expected := cic1ex

	flags := &FlagSet{
		k: 1,
		f: true,
	}

	testOk(in, expected, flags, t)
}

/** InvalidNumbers1 */
var invalidn1in = `89
kh
123
9kjiu
j78`

func TestInvalidNumbers1(t *testing.T) {
	in := invalidn1in

	flags := &FlagSet{
		n: true,
	}

	testFail(in, flags, t)
}

/** InvalidColumns1 */
var invalidc1in = `0 9 3
kj j 3
ki o
oli i 9`

func TestInvalidColumns1(t *testing.T) {
	in := invalidn1in

	flags := &FlagSet{
		k: 2,
	}

	testFail(in, flags, t)
}

/** InvalidColumns2 */
var invalidc2in = `jk j k
iu ui e
io op ww`

func TestInvalidColumns2(t *testing.T) {
	in := invalidc2in

	flags := &FlagSet{
		k: 3,
	}

	testFail(in, flags, t)
}

/** Generic test functions */
func testOk(in, expected string, flags *FlagSet, t *testing.T) {
	out, err := sort(in, flags)
	if err != nil {
		t.Errorf("Test failed: %s\n", err)
	}
	if out != expected {
		t.Errorf("Test failed\n"+
			"In: \n%s\n\n"+
			"Out: \n%s\n\n"+
			"Expected: \n%s\n\n",
			in,
			out,
			expected,
		)
	}
}

func testFail(in string, flags *FlagSet, t *testing.T) {
	_, err := sort(in, flags)
	if err == nil {
		t.Errorf("Test failed: expected error")
	}
}
