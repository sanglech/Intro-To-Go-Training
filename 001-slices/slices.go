package slices

// SumFloats takes a slice of float64, calculates the sum total of all of them, and returns it
// If the input is nil, returns 0.0
func SumFloats(floats []float64) float64 {
	res:=0.0
	if(floats==nil) {
		return res
	}
	for _,val:= range(floats){
		res+=val
	}
	return res
}

// GetIntSlice creates a slice of all integers between first and last (inclusive).
// If first > last GetIntSlice returns an uninitialized slice (or nil).
// For example, GetIntSlice(3, 6) will return: {3, 4, 5, 6}
// GetIntSlice(11, 11) will return: {11}
func GetIntSlice(first, last int) []int {
	var res []int
	if(first>last){
		return nil
	}

	for i := first; i <=last; i++ {
		res = append(res, i)
	}

	return res
}

// ConcatenateStringSlices combines two slices of string into a single slice
// If one of the two slices is empty/nil, the other will be returned.
// If both are nil, returns an uninitialized slice (or nil).
//
// Hint: there is a way to implement this in one line
func ConcatenateStringSlices(sliceA, sliceB []string) []string {
	return append(sliceA,sliceB...)
}
