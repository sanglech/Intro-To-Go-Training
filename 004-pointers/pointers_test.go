package pointers

import (
	"introduction-to-go-training/testutil"
	"math/rand"
	"testing"
)

// Tests for AddToPointer

func TestAddToPointer_InputPtrIsNil_DoesNotPanic(t *testing.T) {
	AddToPointer(nil, 0)

	// if the code panics, this test will fail
}

func TestAddToPointer_AddZero_InputPtrVarIsNotChanged(t *testing.T) {
	input := 0

	AddToPointer(&input, 0)

	testutil.AssertEqual(t, 0, input)
}

func TestAddToPointer_AddPositiveInteger_PtrHasBeenIncreasedByTheIntAmount(t *testing.T) {
	input := 0

	AddToPointer(&input, 5)

	testutil.AssertEqual(t, 5, input)
}

func TestAddToPointer_AddNegativeInteger_PtrHasBeenDecreasedByTheIntAmount(t *testing.T) {
	input := 0

	AddToPointer(&input, -5)

	testutil.AssertEqual(t, -5, input)
}

func TestAddToPointer_CallTwice_BothCallsWorkAsExpected(t *testing.T) {
	input := 0

	AddToPointer(&input, -5)
	testutil.AssertEqual(t, -5, input)

	AddToPointer(&input, 25)
	testutil.AssertEqual(t, 20, input)
}

// Tests for SwapStrings

func TestSwapStrings_BothInputsNil_DoesNotPanic(t *testing.T) {
	SwapStrings(nil, nil)

	// if the code panics, this test will fail
}

func TestSwapStrings_LeftInputNil_DoesNothing(t *testing.T) {
	right := "hello"

	SwapStrings(nil, &right)

	testutil.AssertEqual(t, "hello", right)
}

func TestSwapStrings_RightInputNil_DoesNothing(t *testing.T) {
	left := "hello"

	SwapStrings(&left, nil)

	testutil.AssertEqual(t, "hello", left)
}

func TestSwapStrings_BothInputsPresent_ValuesShouldBeSwapped(t *testing.T) {
	left := "left"
	right := "right"

	SwapStrings(&left, &right)

	testutil.AssertEqual(t, "right", left)
	testutil.AssertEqual(t, "left", right)
}

func TestSwapStrings_DoubleSwap_ValuesShouldBeSwappedBackToTheirOriginalPositions(t *testing.T) {
	left := "left"
	right := "right"

	SwapStrings(&left, &right)
	SwapStrings(&left, &right)

	testutil.AssertEqual(t, "left", left)
	testutil.AssertEqual(t, "right", right)
}

// Tests for SumOptionalList

func TestSumOptionalList_EmptyListInput_ReturnsZero(t *testing.T) {
	res := SumOptionalList(nil)

	testutil.AssertEqual(t, 0, res)
}

func TestSumOptionalList_OnlyNilPtrList_ReturnsZero(t *testing.T) {
	input := []*int{nil, nil, nil}

	res := SumOptionalList(input)

	testutil.AssertEqual(t, 0, res)
}

func TestSumOptionalList_NonNilList_SumsAll(t *testing.T) {
	one := rand.Int()
	two := rand.Int()
	three := rand.Int()
	input := []*int{&one, &two, &three}

	res := SumOptionalList(input)

	testutil.AssertEqual(t, one+two+three, res)
}

func TestSumOptionalList_SomeNils_SumsAllNonNil(t *testing.T) {
	one := rand.Int()
	two := rand.Int()
	input := []*int{nil, &one, nil, &two, nil}

	res := SumOptionalList(input)

	testutil.AssertEqual(t, one+two, res)
}
