package generics

import (
	"introduction-to-go-training/testutil"
	"testing"
)

var testIsInRange_RangeIsValidatedCorrectlyFloat64_cases = []struct {
	Name     string
	Min      float64
	Max      float64
	Value    float64
	Expected bool
}{
	{Name: "equal to min should return true", Min: 3, Max: 10, Value: 3, Expected: true},
	{Name: "equal to max should return true", Min: 3, Max: 10, Value: 10, Expected: true},
	{Name: "between min and max should return true", Min: 3, Max: 5, Value: 4, Expected: true},
	{Name: "equal to min and max should return true", Min: 3, Max: 3, Value: 3, Expected: true},
	{Name: "below min should return false", Min: 3, Max: 5, Value: 2, Expected: false},
	{Name: "above max should return false", Min: 3, Max: 5, Value: 6, Expected: false},
}

func TestIsInRange_RangeIsValidatedCorrectlyFloat64(t *testing.T) {
	for _, testcase := range testIsInRange_RangeIsValidatedCorrectlyFloat64_cases {
		t.Run(testcase.Name, func(t *testing.T) {
			result := IsInRange(testcase.Min, testcase.Max, testcase.Value)

			testutil.AssertEqual(t, testcase.Expected, result)
		})
	}
}

var testIsInRange_RangeIsValidatedCorrectlyInt_cases = []struct {
	Name     string
	Min      int
	Max      int
	Value    int
	Expected bool
}{
	{Name: "equal to min should return true", Min: 3, Max: 10, Value: 3, Expected: true},
	{Name: "equal to max should return true", Min: 3, Max: 10, Value: 10, Expected: true},
	{Name: "between min and max should return true", Min: 3, Max: 5, Value: 4, Expected: true},
	{Name: "equal to min and max should return true", Min: 3, Max: 3, Value: 3, Expected: true},
	{Name: "below min should return false", Min: 3, Max: 5, Value: 2, Expected: false},
	{Name: "above max should return false", Min: 3, Max: 5, Value: 6, Expected: false},
}

func TestIsInRange_RangeIsValidatedCorrectlyInt(t *testing.T) {
	for _, testcase := range testIsInRange_RangeIsValidatedCorrectlyInt_cases {
		t.Run(testcase.Name, func(t *testing.T) {
			result := IsInRange(testcase.Min, testcase.Max, testcase.Value)

			testutil.AssertEqual(t, testcase.Expected, result)
		})
	}
}

var testIsInRange_RangeIsValidatedCorrectlyUint8_cases = []struct {
	Name     string
	Min      uint
	Max      uint
	Value    uint
	Expected bool
}{
	{Name: "equal to min should return true", Min: 3, Max: 10, Value: 3, Expected: true},
	{Name: "equal to max should return true", Min: 3, Max: 10, Value: 10, Expected: true},
	{Name: "between min and max should return true", Min: 3, Max: 5, Value: 4, Expected: true},
	{Name: "equal to min and max should return true", Min: 3, Max: 3, Value: 3, Expected: true},
	{Name: "below min should return false", Min: 3, Max: 5, Value: 2, Expected: false},
	{Name: "above max should return false", Min: 3, Max: 5, Value: 6, Expected: false},
}

func TestIsInRange_RangeIsValidatedCorrectlyUint(t *testing.T) {
	for _, testcase := range testIsInRange_RangeIsValidatedCorrectlyUint8_cases {
		t.Run(testcase.Name, func(t *testing.T) {
			result := IsInRange(testcase.Min, testcase.Max, testcase.Value)

			testutil.AssertEqual(t, testcase.Expected, result)
		})
	}
}

func TestPointerToNillable_InputNil_ReturnsNilNillable(t *testing.T) {
	var nilPtr *int

	result := PointerToNillable(nilPtr)

	testutil.AssertEqual(t, true, result.IsNil)
}

func TestPointerToNillable_InputIntPtr_ReturnsIntValueNillable(t *testing.T) {
	testInt := 12345

	result := PointerToNillable(&testInt)

	testutil.AssertEqual(t, false, result.IsNil)
	testutil.AssertEqual(t, 12345, result.Value)
}

func TestPointerToNillable_InputStructPtr_ReturnsStructValueNillable(t *testing.T) {
	type idStruct struct {
		ID    int
		IDStr string
	}

	testStruct := idStruct{
		ID:    123456,
		IDStr: "testing",
	}

	result := PointerToNillable(&testStruct)

	testutil.AssertEqual(t, false, result.IsNil)
	testutil.AssertEqual(t,
		idStruct{
			ID:    123456,
			IDStr: "testing",
		},
		result.Value)
}

func TestLinkedNode_Next_NoNext_ReturnsFalseForNext(t *testing.T) {
	result := NewLinkedNode(123, nil)

	_, hasNext := result.Next()
	testutil.AssertEqual(t, false, hasNext)
}

func TestLinkedNode_Next_HasNext_ReturnsTrueForNext(t *testing.T) {
	nextNode := NewLinkedNode("test1", nil)
	startNode := NewLinkedNode("test2", nextNode)

	_, hasNext := startNode.Next()

	testutil.AssertEqual(t, true, hasNext)
}

func TestLinkedNode_Value_ChainOfNodes_CanIterateAndGetValueForEach_string(t *testing.T) {
	nodeThree := NewLinkedNode("test3", nil)
	nodeTwo := NewLinkedNode("test2", nodeThree)
	startNode := NewLinkedNode("test1", nodeTwo)

	secondNode, hasSecond := startNode.Next()
	finalNode, hasFinal := secondNode.Next()
	_, finalHasNext := finalNode.Next()

	testutil.AssertEqual(t, "test1", startNode.Value())
	testutil.AssertEqual(t, true, hasSecond)
	testutil.AssertEqual(t, "test2", secondNode.Value())
	testutil.AssertEqual(t, true, hasFinal)
	testutil.AssertEqual(t, "test3", finalNode.Value())
	testutil.AssertEqual(t, false, finalHasNext)
}

func TestLinkedNode_Value_ChainOfNodes_CanIterateAndGetValueForEach_float64(t *testing.T) {
	nodeThree := NewLinkedNode(3.3, nil)
	nodeTwo := NewLinkedNode(2.2, nodeThree)
	startNode := NewLinkedNode(1.1, nodeTwo)

	secondNode, hasSecond := startNode.Next()
	finalNode, hasFinal := secondNode.Next()
	_, finalHasNext := finalNode.Next()

	testutil.AssertEqual(t, 1.1, startNode.Value())
	testutil.AssertEqual(t, true, hasSecond)
	testutil.AssertEqual(t, 2.2, secondNode.Value())
	testutil.AssertEqual(t, true, hasFinal)
	testutil.AssertEqual(t, 3.3, finalNode.Value())
	testutil.AssertEqual(t, false, finalHasNext)
}
