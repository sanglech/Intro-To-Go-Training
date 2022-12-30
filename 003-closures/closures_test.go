package closures

import (
	"introduction-to-go-training/testutil"
	"strings"
	"testing"
)

// Tests for GetAdderSubber

func TestGetAdderSubber_InitialStateIsZero(t *testing.T) {
	_, _, curval := GetAdderSubber()

	testutil.AssertEqual(t, 0, curval())
}

func TestGetAdderSubber_InitialStateCanBeModifiedWithAdderAndSubber(t *testing.T) {
	addr, subbr, curval := GetAdderSubber()

	addr(10)
	testutil.AssertEqual(t, 10, curval())
	addr(1)
	testutil.AssertEqual(t, 11, curval())
	addr(0)
	testutil.AssertEqual(t, 11, curval())
	addr(-1)
	testutil.AssertEqual(t, 10, curval())

	subbr(0)
	testutil.AssertEqual(t, 10, curval())
	subbr(5)
	testutil.AssertEqual(t, 5, curval())
	subbr(1)
	testutil.AssertEqual(t, 4, curval())
	subbr(-22)
	testutil.AssertEqual(t, 26, curval())
}

// Tests for NormalizeStrReader

func TestNormalizeStrReader_upperTrue_shouldConvertAllOutputToUppercase(t *testing.T) {
	ctr := 0
	sr := func() string {
		ctr++
		return strings.Repeat("aA", ctr)
	}

	nsr := NormalizeStrReader(sr, true)

	testutil.AssertEqual(t, "AA", nsr())
	testutil.AssertEqual(t, "AAAA", nsr())
	testutil.AssertEqual(t, "AAAAAA", nsr())
	testutil.AssertEqual(t, "AAAAAAAA", nsr())
}

func TestNormalizeStrReader_upperFalse_shouldConvertAllOutputToLowercase(t *testing.T) {
	ctr := 0
	sr := func() string {
		ctr++
		return strings.Repeat("Bb", ctr)
	}

	nsr := NormalizeStrReader(sr, false)

	testutil.AssertEqual(t, "bb", nsr())
	testutil.AssertEqual(t, "bbbb", nsr())
	testutil.AssertEqual(t, "bbbbbb", nsr())
	testutil.AssertEqual(t, "bbbbbbbb", nsr())
}

func TestNormalizeStrReader_strReaderReturnsEmptyStr_shouldReturnEmptyStr(t *testing.T) {
	sr := func() string {
		return ""
	}

	nsr := NormalizeStrReader(sr, false)

	testutil.AssertEqual(t, "", nsr())
	testutil.AssertEqual(t, "", nsr())
}

func TestNormalizeStrReader_strReaderReturnsNumbers_shouldNotChangeStrReaderOutput(t *testing.T) {
	sr := func() string {
		return "123.67&"
	}

	nsr := NormalizeStrReader(sr, true)

	testutil.AssertEqual(t, "123.67&", nsr())
	testutil.AssertEqual(t, "123.67&", nsr())
}
