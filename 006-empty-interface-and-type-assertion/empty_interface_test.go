package emptyint

import (
	"introduction-to-go-training/testutil"
	"testing"
)

// Tests for MultitypeSum

func TestMultitypeSum_EmptyListInput_ReturnsZero(t *testing.T) {
	res := MultitypeSum([]interface{}{})

	testutil.AssertEqual(t, 0.0, res)
}

func TestMultitypeSum_AllSupportedTypes_ReturnsSumOfAllInput(t *testing.T) {
	res := MultitypeSum([]interface{}{0.11, uint64(1), int64(2)})

	exptd := 0.11 + float64(uint64(1)) + float64(int64(2))
	testutil.AssertEqual(t, exptd, res)
}

func TestMultitypeSum_IncludingUnsupportedTypes_ReturnsSumExcludingUnsupported(t *testing.T) {
	res := MultitypeSum([]interface{}{0.11, []int{1}, float32(0.22), uint64(1), int64(2), "dfhsdf"})

	exptd := 0.11 + float64(uint64(1)) + float64(int64(2))
	testutil.AssertEqual(t, exptd, res)
}

func TestMultitypeSum_AllUnsupportedTypes_ReturnsZero(t *testing.T) {
	res := MultitypeSum([]interface{}{[]int{1}, float32(0.22), "dfhsdf"})

	testutil.AssertEqual(t, 0.0, res)
}

// tests for AppendIfStringer

func TestAppendIfStringer_nilInput_returnsJustTheString(t *testing.T) {
	inputStr := ": appended string"

	res := AppendIfStringer(nil, inputStr)

	if res == nil {
		t.Fatal("should not return nil")
	}
	testutil.AssertEqual(t, inputStr, res.String())
}

func TestAppendIfStringer_notImplOfStringer_returnsJustTheString(t *testing.T) {
	inputStruct := struct{}{}
	inputStr := ": appended string"

	res := AppendIfStringer(inputStruct, inputStr)

	if res == nil {
		t.Fatal("should not return nil")
	}
	testutil.AssertEqual(t, inputStr, res.String())
}

type testStringerImpl struct {
	Str string
}

func (t testStringerImpl) String() string {
	return t.Str
}

func TestAppendIfStringer_stringerInput_returnsConcatenatedStringer(t *testing.T) {
	inputStruct := testStringerImpl{Str: "stringerstring"}
	inputStr := ": appended string"

	res := AppendIfStringer(inputStruct, inputStr)

	if res == nil {
		t.Fatal("should not return nil")
	}
	testutil.AssertEqual(t, "stringerstring: appended string", res.String())
}
