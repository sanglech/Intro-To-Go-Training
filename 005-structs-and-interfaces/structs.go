package structs

// AdderSubber maintains an internal sum which can be modified using the interface's methods
// Do not modify this interface!
type AdderSubber interface {

	// Add an integer to the internal sum
	Add(amount int)

	// Subtract an integer from the internal sum
	Subtract(amount int)

	// GetCurrentValue returns the current internal sum
	GetCurrentValue() int
}

type SumInt struct{
	sum int
}

// NewAdderSubber returns a concrete implementation of the AdderSubber interface.
// The returned AdderSubber will have its sum set to initialSum
// Note: feel free to define any new types you need to achieve this.
//		 And remember that receiver arguments (like all Go arguments) are pass-by-value.
func NewAdderSubber(initialSum int) AdderSubber {
	return &SumInt{sum:initialSum}
}

func (s *SumInt) Add(amount int){
	s.sum+=amount
}

func (s *SumInt) Subtract(amount int){
	s.sum-=amount
}
func (s *SumInt) GetCurrentValue() int{
	return s.sum
}


// Dog is an interface which can be satisfied by an individual of the species Canis Familiaris.
// Do not modify this interface!
type Dog interface {

	// MakeNoise returns a dog barking noise.
	MakeNoise() string

	// RollOver returns a boolean indicating whether the dog rolled over as instructed.
	// If the dog is a good dog it will roll over. If it is not a good dog, it will not.
	RollOver() bool

	// SetIsGoodDog sets a flag indicating whether this Dog is a good dog.
	SetIsGoodDog(isGoodDog bool)
}

// Barker is an animal that makes a barking noise.
//
// Do not modify this struct!
type Barker struct{}

// MakeNoise returns the noise that barker makes.
//
// Do not modify this Barker method!
func (n *Barker) MakeNoise() string {
	return "BARK BARK!!!!"
}

// NewDog returns a concrete struct which meets the Dog interface (see above).
//
// To solve this problem, you should use the Barker struct in your Dog implementation.
// Hint: remember that we talked about embedded structs.
//
// It doesn't matter whether your impl is or is not a good dog by default.
type Bahaman struct {
	Barker
	isGoodDog bool
}

func (d *Bahaman) RollOver () bool{
	return d.isGoodDog;
}

func (d *Bahaman) SetIsGoodDog(isGoodDog bool){
	d.isGoodDog = isGoodDog
}

func NewDog() Dog {
	return &Bahaman{isGoodDog:true}
}
