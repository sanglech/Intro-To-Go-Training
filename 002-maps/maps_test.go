package maps

import (
	"introduction-to-go-training/testutil"
	"sort"
	"testing"
)

// tests for GetKeyValSlice

func TestGetKeyValSlice_NilMap_ReturnsEmptySlice(t *testing.T) {
	res := GetKeyValSlice(nil)

	testutil.AssertEqualMsgf(t, 0, len(res), "Expected length %v but got length %v")
}

func TestGetKeyValSlice_OneKey_ConcatenatesAndReturnsTheKey(t *testing.T) {
	res := GetKeyValSlice(map[string]string{"abc": "123"})

	expected := []string{"abc123"}
	testutil.AssertEqualMsgf(t, len(expected), len(res), "Expected length %v but got length %v")
	for itr, val := range res {
		testutil.AssertEqualMsgf(t, expected[itr], val, "Expected %v but got %v (slice index %v)", itr)
	}
}

func TestGetKeyValSlice_SeveralKeys_ConcatenatesAndReturnsAllKeys(t *testing.T) {
	res := GetKeyValSlice(map[string]string{"abc": "123", "def": "456", "ghi": "789"})

	expected := []string{"abc123", "def456", "ghi789"}
	testutil.AssertEqualMsgf(t, len(expected), len(res), "Expected length %v but got length %v")
	sort.Strings(expected)
	sort.Strings(res)
	for itr, val := range res {
		testutil.AssertEqualMsgf(t, expected[itr], val, "Expected %v but got %v (slice index %v)", itr)
	}
}

// tests for DeleteFromMap

func TestDeleteFromMap_Nil_DoesntPanic(t *testing.T) {
	DeleteFromMap(nil, []int{})

	// if there is no panic, this test will pass
}

func TestDeleteFromMap_EmptyMap_DoesNothing(t *testing.T) {
	var input map[int]int
	DeleteFromMap(input, []int{})

	testutil.AssertEqualMsgf(t, 0, len(input), "Expected length %v but got length %v")
}

func TestDeleteFromMap_DeleteKeyNotInMap_DoesNothing(t *testing.T) {
	input := map[int]int{1: 100, 2: 200}
	DeleteFromMap(input, []int{3})

	testutil.AssertEqualMsgf(t, 2, len(input), "Expected length %v but got length %v")
	testutil.AssertEqual(t, 100, input[1])
	testutil.AssertEqual(t, 200, input[2])
}

func TestDeleteFromMap_DeleteSomeKeys_AllKeysAreRemoved(t *testing.T) {
	input := map[int]int{1: 100, 2: 200, 3: 300}
	DeleteFromMap(input, []int{1, 2})

	testutil.AssertEqualMsgf(t, 1, len(input), "Expected length %v but got length %v")
	testutil.AssertEqual(t, 300, input[3])
}

// tests for SetAndGet

func TestSetAndGet_NilMap_ReturnsEmptyStr(t *testing.T) {
	res := SetAndGet(nil, 1.1, "abc", 2.2)

	testutil.AssertEqual(t, "", res)
}

func TestSetAndGet_SetKeyAndGetDoesNotExist_KeyIsSetAndReturnsNOTEXIST(t *testing.T) {
	input := make(map[float64]string)
	res := SetAndGet(input, 1.1, "abc", 2.2)

	testutil.AssertEqualMsgf(t, 1, len(input), "Expected length %v but got length %v")
	testutil.AssertEqual(t, "NOTEXIST", res)
	testutil.AssertEqual(t, "abc", input[1.1])
}

func TestSetAndGet_SetOverwritesAndGetDoesExist_KeyIsSetAndGetReturnsValue(t *testing.T) {
	input := map[float64]string{
		1.1: "def",
		2.2: "zyx",
	}
	res := SetAndGet(input, 1.1, "abc", 2.2)

	testutil.AssertEqualMsgf(t, 2, len(input), "Expected length %v but got length %v")
	testutil.AssertEqual(t, "zyx", res)
	testutil.AssertEqual(t, "abc", input[1.1])
}

func TestSetAndGet_SetAndGetOnSameKey_ValueThatWasSetIsGotten(t *testing.T) {
	input := make(map[float64]string)
	res := SetAndGet(input, 1.1, "abc", 1.1)

	testutil.AssertEqualMsgf(t, 1, len(input), "Expected length %v but got length %v")
	testutil.AssertEqual(t, "abc", res)
	testutil.AssertEqual(t, "abc", input[1.1])
}
