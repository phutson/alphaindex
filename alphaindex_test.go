package alphaindex_test

import (
	"fmt"
	"testing"

	ain "github.com/phutson/alphaindex"
)

func TestNew(t *testing.T) {
	_, err := ain.NewIndex("aa1w")
	if err == nil {
		fmt.Println("did not error on an invalid index")
		t.FailNow()
	}

	CCI, err := ain.NewIndex("AaBb")
	if err != nil {
		fmt.Println("error was returned when there should not have been one ", err)
		t.FailNow()
	}

	if CCI.GetIndexString() != "AABB" {
		fmt.Println("did not convert to upper case")
		t.FailNow()
	}
}

func TestInc(t *testing.T) {
	CCI, err := ain.NewIndex("")
	if err != nil {
		fmt.Println("should not have errored ", err)
		t.FailNow()
	}
	CCI.Increment()
	if CCI.GetIndexString() != "A" {
		fmt.Println("Was not A ", CCI.GetIndexString())
		t.FailNow()
	}
}

func TestInc2(t *testing.T) {
	CCI, err := ain.NewIndex("AA")
	if err != nil {
		fmt.Println("should not have errored ", err)
		t.FailNow()
	}
	CCI.Increment()
	if CCI.GetIndexString() != "AB" {
		fmt.Println("Was not AB ", CCI.GetIndexString())
		t.FailNow()
	}
}

func TestInc3(t *testing.T) {
	CCI, err := ain.NewIndex("AZ")
	if err != nil {
		fmt.Println("should not have errored ", err)
		t.FailNow()
	}
	CCI.Increment()
	if CCI.GetIndexString() != "BA" {
		fmt.Println("Was not BA ", CCI.GetIndexString())
		t.FailNow()
	}
}

func TestDec(t *testing.T) {
	AIN, err := ain.NewIndex("AZ")
	if err != nil {
		fmt.Println("should not have errored ", err)
		t.FailNow()
	}
	AIN.Decriment()
	if AIN.GetIndexString() !=
		"AY" {
		fmt.Println("Was not AY ", AIN.GetIndexString())
		t.FailNow()
	}
}

func TestDec2(t *testing.T) {
	AIN, err := ain.NewIndex("AA")
	if err != nil {
		fmt.Println("should not have errored ", err)
		t.FailNow()
	}
	AIN.Decriment()
	if AIN.GetIndexString() != "Z" {
		fmt.Println("Was not Z ", AIN.GetIndexString())
		t.FailNow()
	}
}

func TestRecursiveInc(t *testing.T) {
	index := []byte{'A'}
	outIndex := ain.IncRecursive(index, 0)
	if len(outIndex) != 1 {
		fmt.Println("out index was not 1 character")
		t.FailNow()
	}
	if outIndex[0] != 'B' {
		fmt.Println("Was not incremented correctly ", outIndex[0])
		t.FailNow()
	}
}

func TestRecursiveInc2(t *testing.T) {
	index := []byte{'A', 'Z'}
	outIndex := ain.IncRecursive(index, 1)
	if len(outIndex) != 2 {
		fmt.Println("out index was not 2 characters")
		t.FailNow()
	}
	if outIndex[0] != 'B' {
		fmt.Println("Was not incremented correctly was not b ", string(outIndex[:]))
		t.FailNow()
	}
	if outIndex[1] != 'A' {
		fmt.Println("Was not incremented correctly was not a ", string(outIndex[:]))
		t.FailNow()
	}
}

func TestRecursiveInc3(t *testing.T) {
	index := []byte{'Z'}
	outIndex := ain.IncRecursive(index, 0)
	if len(outIndex) != 2 {
		fmt.Println("out index was not 2 characters")
		t.FailNow()
	}
	if outIndex[0] != 'A' {
		fmt.Println("Was not incremented correctly was not a ", string(outIndex[:]))
		t.FailNow()
	}
	if outIndex[1] != 'A' {
		fmt.Println("Was not incremented correctly was not a ", string(outIndex[:]))
		t.FailNow()
	}
}

func TestRecursiveInc4(t *testing.T) {
	index := []byte{'Z', 'Z', 'Z'}
	outIndex := ain.IncRecursive(index, 2)
	if len(outIndex) != 4 {
		fmt.Println("out index was not 4 characters")
		t.FailNow()
	}
	if outIndex[0] != 'A' {
		fmt.Println("Was not incremented correctly was not a ", string(outIndex[:]))
		t.FailNow()
	}
	if outIndex[3] != 'A' {
		fmt.Println("Was not incremented correctly was not a ", string(outIndex[:]))
		t.FailNow()
	}
}

func TestRecursiveInc5(t *testing.T) {
	index := []byte{'Z', 'B', 'Z'}
	outIndex := ain.IncRecursive(index, 2)
	if len(outIndex) != 3 {
		fmt.Println("out index was not 3 characters")
		t.FailNow()
	}
	if outIndex[0] != 'Z' {
		fmt.Println("Was not incremented correctly was not a ", string(outIndex[:]))
		t.FailNow()
	}
	if outIndex[2] != 'A' {
		fmt.Println("Was not incremented correctly was not a ", string(outIndex[:]))
		t.FailNow()
	}
	if outIndex[1] != 'C' {
		fmt.Println("Was not incremented correctly was not c ", string(outIndex[:]))
		t.FailNow()
	}
}

func TestRecursiveDec(t *testing.T) {
	index := []byte{'Z'}
	outIndex := ain.DecRecursive(index, 0)
	if len(outIndex) != 1 {
		fmt.Println("out index was not 1 character")
		t.FailNow()
	}
	if outIndex[0] != 'Y' {
		fmt.Println("Was not incremented correctly ", outIndex[0])
		t.FailNow()
	}
}

func TestRecursiveDec2(t *testing.T) {
	index := []byte{'A', 'A'}
	index = ain.DecRecursive(index, 1)
	if len(index) != 1 {
		fmt.Println("out index was not 1 characters", len(index), string(index[:]))
		t.FailNow()
	}
	if index[0] != 'Z' {
		fmt.Println("Was not incremented correctly was not b ", string(index[:]))
		t.FailNow()
	}
}

func TestRecursiveDec3(t *testing.T) {
	index := []byte{'Z', 'A', 'A'}
	index = ain.DecRecursive(index, 2)
	if len(index) != 3 {
		fmt.Println("out index was not 1 characters", len(index), string(index[:]))
		t.FailNow()
	}
	if string(index[:]) != "YZZ" {
		fmt.Println("Was not incremented correctly was not YZZ ", string(index[:]))
		t.FailNow()
	}
}
