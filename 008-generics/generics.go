package generics

// RangeCheckable is an interface matching all types which are usable
// with the generic IsInRange function.
type RangeCheckable interface {
	int | float64 | uint
}

// IsInRange checks that the input "value" is between (inclusive) "min" and "max."
// I.e. it checks that "value" is in the interval ["min", "max"] .
//
// Your implementation of this function should be able to support: int, float64, & uint
//
// TODO Important Note! You will need to change the type constraint interface (RangeCheckable) above
// to match the required types.
func IsInRange[T RangeCheckable](min T, max T, value T) bool {

	return value >=min && value <=max

}

// Nillable represents a nillable value in a way that prevents nil pointer errors.
// DO NOT MODIFY THIS STRUCT
type Nillable[T any] struct {

	// IsNil will be true if this Nillable does not contain useful data in its "Value" field.
	IsNil bool

	// Value contains either the zero value for type T (if IsNil is true)
	// or any value of type T (if IsNil is false).
	Value T
}

// PointerToNillable converts the input "value" from a pointer to a Nillable.
// If the pointer is nil, the Nillable should also be "nil."
// And if the pointer is not nil, the Nillable should contain its value.
func PointerToNillable[T any](value *T) Nillable[T] {
	if(value==nil) {
		return Nillable[T]{
			IsNil: true,
		}
	}
	return Nillable[T]{
		IsNil: false,
		Value: *value,
	}

}

// LinkedNode is a data structure which represents one node in a singly linked list.
// It contains a value of type T and a function for getting the next node in the list (if any).
// DO NOT MODIFY THIS INTERFACE
type LinkedNode[T any] interface {

	// Next returns the next LinkedNode in this linked list.
	// If there is a next node, the boolean result will be true and the node will be returned.
	// If there are no more nodes, the boolean result will be false & nil will be returned.
	Next() (LinkedNode[T], bool)

	// Value returns the data value of type T stored in this list node.
	Value() T
}

type LinkedNodeImpl [T any] struct  {
	next  LinkedNode[T]
	value T
}

func (n *LinkedNodeImpl[T]) Next() (LinkedNode[T], bool){
	if (n.next==nil){
		return nil, false;
	}
	return n.next,true;
}

func (n *LinkedNodeImpl[T]) Value() T {
	return n.value;
}

// NewLinkedNode creates a new LinkedNode[T] for use in a linked list.
// It takes input "value" to be stored in the returned node.
// It also takes "next" which is the next LinkedNode[T] in the list after this one.
//
// TODO write your own implementation of the "LinkedNode[T]" interface. And use it to implement this function.
func NewLinkedNode[T any](value T, next LinkedNode[T]) LinkedNode[T] {

	return &LinkedNodeImpl[T]{
		next:  next,
		value: value,
	}

}
