package main

import (
	"testing"
)

/** Ok tests */
func TestPlus1(t *testing.T) {
	in := "2 + 2"
	expected := 4
	testOk(in, expected, t)
}

func TestPlus2(t *testing.T) {
	in := "7 + 9"
	expected := 16
	testOk(in, expected, t)
}

func TestMinus1(t *testing.T) {
	in := "5 - 9"
	expected := -4
	testOk(in, expected, t)
}

func TestMinus2(t *testing.T) {
	in := "89 - 12"
	expected := 77
	testOk(in, expected, t)
}

func TestMultiply1(t *testing.T) {
	in := "9 * 24"
	expected := 216
	testOk(in, expected, t)
}

func TestMultiply2(t *testing.T) {
	in := "3 * 4"
	expected := 12
	testOk(in, expected, t)
}

func TestDivide1(t *testing.T) {
	in := "700 / 35"
	expected := 20
	testOk(in, expected, t)
}

func TestDivide2(t *testing.T) {
	in := "1024 / 8"
	expected := 128
	testOk(in, expected, t)
}

func TestPriority1(t *testing.T) {
	in := "2 + 2 * 2"
	expected := 6
	testOk(in, expected, t)
}

func TestPriority2(t *testing.T) {
	in := "8 - 9 / 3"
	expected := 5
	testOk(in, expected, t)
}

func TestBrace1(t *testing.T) {
	in := "7 * (5 + 4)"
	expected := 63
	testOk(in, expected, t)
}

func TestBrace2(t *testing.T) {
	in := "(7 * (5 + 4) + 2) / 5"
	expected := 13
	testOk(in, expected, t)
}

func TestComplicated(t *testing.T) {
	in := "(2 * 9 / (1 + 2) - 4) * 3 - 2"
	expected := 4
	testOk(in, expected, t)
}

/** Fail tests */
func TestInvalidBrace1(t *testing.T) {
	in := "(23 * 8))"
	testFail(in, t)
}

func TestInvalidBrace2(t *testing.T) {
	in := "45 - ((67 * 9)"
	testFail(in, t)
}

func TestInvalidBrace3(t *testing.T) {
	in := "((86 + 9) * (42 / 8) - (5)"
	testFail(in, t)
}

func TestInvalidOperator1(t *testing.T) {
	in := "45 * * 3"
	testFail(in, t)
}

func TestInvalidOperator2(t *testing.T) {
	in := "(7 - 9) // 2"
	testFail(in, t)
}

func TestInvalidOperator3(t *testing.T) {
	in := "8 +* 4"
	testFail(in, t)
}

/** Generic test functions */
func testOk(in string, expected int, t *testing.T) {
	out, err := calculate(in)
	if err != nil {
		t.Errorf("Test failed: %s\n", err)
	}
	if out != expected {
		t.Errorf("Test failed\n"+
			"In: %s\n"+
			"Out: %d\n"+
			"Expected: %d\n",
			in,
			out,
			expected,
		)
	}
}

func testFail(in string, t *testing.T) {
	_, err := calculate(in)
	if err == nil {
		t.Errorf("Test failed\n"+
			"In: %s\n"+
			"Expected: error\n",
			in,
		)
	}
}
