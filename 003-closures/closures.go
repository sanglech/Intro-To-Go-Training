package closures

import (
	"strings"
)

// GetAdderSubber returns three functions which share the same state: a sum (initialized to zero).
// The first two returned functions should take integer input:
// 		The adder function should add that input to the internal state.
//		The subber function should subtract that input from the internal state.
// The third function just returns the current value of the internal sum.
// E.g.
//      adder, subber, curVal := GetAdderSubber()
// 		adder(5) // internal sum is 5
// 		adder(11) // internal sum is 16
// 		subber(6) // internal sum  is 10
// 		curVal() // returns 10
func GetAdderSubber() (adder func(int), subber func(int), curVal func() int) {
	sum:=0
	adder=func(toAdd int){sum+=toAdd}
	subber=func(toSub int) {sum-=toSub}
	curVal=func() int {return sum}

	return adder, subber, curVal
}

// NormalizeStrReader takes a function and a boolean as input.
// The input function (strReader) returns a string every time it is called.
// NormalizeStrReader should return a function which wraps strReader and "normalizes" its output.
// The output should be normalized by making it either all UPPERCASE or all lowercase.
// If the "upper" parameter is true, it should be all UPPERCASE. If "upper" is false, all lowercase.
// E.g.
//		sr := func() string { return "TeSt OuTpUt" }
//
// 		nsr := NormalizeStrReader(sr, true)
//      nsr() // returns "TEST OUTPUT"
//
//     nsrLower := NormalizeStrReader(sr, false)
//     nsrLower() // returns "test output"
//
// Hint: the "strings" standard library package contains ToUpper and ToLower functions.
func NormalizeStrReader(strReader func() string, upper bool) func() string {
	return func() string{
		if(upper){
			return strings.ToUpper(strReader())
		} else {
			return strings.ToLower(strReader())
		}
	}
}
