package emptyint

// MultitypeSum adds a list of different types of numbers together, producing a float64.
// It can handle the following types:
//			float64, int64, uint64
// If the input list contains any other type, that element of the list is ignored.
// If the list is empty (or consists of only non-supported types) returns 0.0
func MultitypeSum(input []interface{}) float64 {
	res :=0.0

	for _, val := range input {
		switch actualVal := val.(type) {
		case int64:
			res += float64(actualVal)
		case float64:
			res += actualVal
		case uint64:
			res+=float64(actualVal)
		}
	}
	return res 
}

// Stringer is an interface for something which can be represented as a string
// Do not modify this interface!
type Stringer interface {

	// String gets the string representation of this Stringer
	String() string
}

type AStringer struct{
	str string
}

// AppendIfStringer checks if "input" is an implementation of Stringer and, if yes, returns a new Stringer
// with the value of "str" appended to the output of input.String():
//		inputIfStringer.String() + str
// If "input" is not a Stringer, returns a new Stringer containing only the value of "str".

func (s *AStringer) String() string{
	return s.str;
}

func AppendIfStringer(input interface{}, str string) Stringer {
	val,isStringer:= input.(Stringer)

	if(isStringer){
		return &AStringer{str: val.String() + str}
	}
	return &AStringer{str}
}
