package main

import (
	"reflect"
	"testing"
)

/** Sort1 */
func TestSort1(t *testing.T) {
	in := []string{
		"Soda",
		"Apple",
		"Join",
		"Unite",
		"Book",
		"Debian",
		"BOOK",
	}
	expected := []string{
		"Apple",
		"BOOK",
		"Book",
		"Debian",
		"Join",
		"Soda",
		"Unite",
	}

	flags := &FlagSet{}

	testOk(in, expected, flags, t)
}

/** SortReverse1 */
func TestSortReverse1(t *testing.T) {
	in := []string{
		"Soda",
		"Apple",
		"Join",
		"Unite",
		"Book",
		"Debian",
		"BOOK",
	}
	expected := []string{
		"Unite",
		"Soda",
		"Join",
		"Debian",
		"Book",
		"BOOK",
		"Apple",
	}

	flags := &FlagSet{
		reverse: true,
	}

	testOk(in, expected, flags, t)
}

/** Sort2 */
func TestSort2(t *testing.T) {
	in := []string{
		"7Up",
		"1Helloworld",
		"4From",
		"1Looser",
		"0Urgent",
	}
	expected := []string{
		"0Urgent",
		"1Helloworld",
		"1Looser",
		"4From",
		"7Up",
	}

	flags := &FlagSet{}

	testOk(in, expected, flags, t)
}

/** IgnoreCase1 */
func TestIgnoreCase1(t *testing.T) {
	in := []string{
		"Join",
		"jOhn",
		"ko13",
		"KO09",
		"Lolipop",
	}
	expected := []string{
		"jOhn",
		"Join",
		"KO09",
		"ko13",
		"Lolipop",
	}

	flags := &FlagSet{
		ignoreCase: true,
	}

	testOk(in, expected, flags, t)
}

/** IgnoreCaseReverse1 */
func TestIgnoreCaseReverse1(t *testing.T) {
	in := []string{
		"Join",
		"jOhn",
		"ko13",
		"KO09",
		"Lolipop",
	}
	expected := []string{
		"Lolipop",
		"ko13",
		"KO09",
		"Join",
		"jOhn",
	}

	flags := &FlagSet{
		ignoreCase: true,
		reverse:    true,
	}

	testOk(in, expected, flags, t)
}

/** IgnoreCaseUnique1 */
func TestIgnoreCaseUnique1(t *testing.T) {
	in := []string{
		"JOHN",
		"Join",
		"jOhn",
		"ko13",
		"KO09",
		"Lolipop",
	}
	expected := []string{
		"JOHN",
		"Join",
		"KO09",
		"ko13",
		"Lolipop",
	}

	flags := &FlagSet{
		ignoreCase: true,
		unique:     true,
	}

	testOk(in, expected, flags, t)
}

/** IgnoreCaseUniqueReverse1 */
func TestIgnoreCaseUniqueReverse1(t *testing.T) {
	in := []string{
		"Join",
		"jOhn",
		"ko13",
		"JOHN",
		"KO09",
		"Lolipop",
	}
	expected := []string{
		"Lolipop",
		"ko13",
		"KO09",
		"Join",
		"JOHN",
	}

	flags := &FlagSet{
		ignoreCase: true,
		unique:     true,
		reverse:    true,
	}

	testOk(in, expected, flags, t)
}

/** Unique1 */
func TestUnique1(t *testing.T) {
	in := []string{
		"Lol",
		"Bom",
		"Kohn",
		"Bom",
		"Tideman",
	}
	expected := []string{
		"Bom",
		"Kohn",
		"Lol",
		"Tideman",
	}

	flags := &FlagSet{
		unique: true,
	}

	testOk(in, expected, flags, t)
}

/** UniqueReverse1 */
func TestUniqueReverse1(t *testing.T) {
	in := []string{
		"Lol",
		"Bom",
		"Kohn",
		"Bom",
		"Tideman",
	}
	expected := []string{
		"Tideman",
		"Lol",
		"Kohn",
		"Bom",
	}

	flags := &FlagSet{
		unique:  true,
		reverse: true,
	}

	testOk(in, expected, flags, t)
}

/** Numbers1 */
func TestNumbers1(t *testing.T) {
	in := []string{
		"9",
		"76",
		"12",
		"67",
		"994",
		"-234",
	}
	expected := []string{
		"-234",
		"9",
		"12",
		"67",
		"76",
		"994",
	}

	flags := &FlagSet{
		numbers: true,
	}

	testOk(in, expected, flags, t)
}

/** NumbersReverse1 */
func TestNumbersReverse1(t *testing.T) {
	in := []string{
		"9",
		"76",
		"12",
		"67",
		"994",
		"-234",
	}
	expected := []string{
		"994",
		"76",
		"67",
		"12",
		"9",
		"-234",
	}

	flags := &FlagSet{
		numbers: true,
		reverse: true,
	}

	testOk(in, expected, flags, t)
}

/** Columns1 */
func TestColumns1(t *testing.T) {
	in := []string{
		"Jon Snow",
		"Jorah Marmont",
		"Deyneris Targarien",
		"Aria Stark",
	}
	expected := []string{
		"Jorah Marmont",
		"Jon Snow",
		"Aria Stark",
		"Deyneris Targarien",
	}

	flags := &FlagSet{
		column: 1,
	}

	testOk(in, expected, flags, t)
}

/** ColumnsNumbers1 */
func TestColumnsNumbers1(t *testing.T) {
	in := []string{
		"923 8 122",
		"87 23 87",
		"-98 23 9234844",
		"221 764 -23",
		"8765 -90 44",
		"23123 -123 987",
	}
	expected := []string{
		"221 764 -23",
		"8765 -90 44",
		"87 23 87",
		"923 8 122",
		"23123 -123 987",
		"-98 23 9234844",
	}

	flags := &FlagSet{
		column:  2,
		numbers: true,
	}

	testOk(in, expected, flags, t)
}

/** ColumnsIgnoreCase1 */
func TestColumnsIgnoreCase1(t *testing.T) {
	in := []string{
		"1 Kok",
		"1 KOj",
		"1 kOa",
		"1 koh",
		"1 koB",
	}
	expected := []string{
		"1 kOa",
		"1 koB",
		"1 koh",
		"1 KOj",
		"1 Kok",
	}

	flags := &FlagSet{
		column:     1,
		ignoreCase: true,
	}

	testOk(in, expected, flags, t)
}

/** InvalidNumbers1 */
func TestInvalidNumbers1(t *testing.T) {
	in := []string{
		"89",
		"kh",
		"123",
		"9kjiu",
		"j78",
	}

	flags := &FlagSet{
		numbers: true,
	}

	testFail(in, flags, t)
}

/** InvalidColumns1 */
func TestInvalidColumns1(t *testing.T) {
	in := []string{
		"0 9 3",
		"kj j 3",
		"ki o",
		"oli i 9",
	}

	flags := &FlagSet{
		column: 2,
	}

	testFail(in, flags, t)
}

/** InvalidColumns2 */
func TestInvalidColumns2(t *testing.T) {
	in := []string{
		"jk j k",
		"iu ui e",
		"io op ww",
	}

	flags := &FlagSet{
		column: 3,
	}

	testFail(in, flags, t)
}

/** Generic test functions */
func testOk(in, expected []string, flags *FlagSet, t *testing.T) {
	out, err := sort(append([]string(nil), in...), flags)
	if err != nil {
		t.Errorf("Test failed: %s\n", err)
	}
	if !reflect.DeepEqual(out, expected) {
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

func testFail(in []string, flags *FlagSet, t *testing.T) {
	_, err := sort(in, flags)
	if err == nil {
		t.Errorf("Test failed: expected error")
	}
}
