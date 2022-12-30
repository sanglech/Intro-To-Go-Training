package structs

import (
	"introduction-to-go-training/testutil"
	"reflect"
	"testing"
)

// Tests for NewAdderSubber

func TestNewAdderSubber_NoAddsOrSubs_ShouldHaveThePassedInitValue(t *testing.T) {
	as := NewAdderSubber(123)

	testutil.AssertEqual(t, 123, as.GetCurrentValue())
}

func TestNewAdderSubber_Add_ShouldIncreaseSumByTheAddedValue(t *testing.T) {
	as := NewAdderSubber(0)

	testutil.AssertEqual(t, 0, as.GetCurrentValue())
	as.Add(56)
	testutil.AssertEqual(t, 56, as.GetCurrentValue())
	as.Add(22)
	testutil.AssertEqual(t, 78, as.GetCurrentValue())
	as.Add(-79)
	testutil.AssertEqual(t, -1, as.GetCurrentValue())
}

func TestNewAdderSubber_Subtract_ShouldDecreaseSumByTheSubtractedValue(t *testing.T) {
	as := NewAdderSubber(100)

	testutil.AssertEqual(t, 100, as.GetCurrentValue())
	as.Subtract(56)
	testutil.AssertEqual(t, 44, as.GetCurrentValue())
	as.Subtract(22)
	testutil.AssertEqual(t, 22, as.GetCurrentValue())
	as.Subtract(-44)
	testutil.AssertEqual(t, 66, as.GetCurrentValue())
}

func TestNewAdderSubber_AddAndSubtract_ShouldAddAndSubtractFromTheSumAsExpected(t *testing.T) {
	as := NewAdderSubber(10000)

	testutil.AssertEqual(t, 10000, as.GetCurrentValue())
	as.Subtract(100)
	testutil.AssertEqual(t, 9900, as.GetCurrentValue())
	as.Add(101)
	testutil.AssertEqual(t, 10001, as.GetCurrentValue())
	as.Subtract(-1)
	testutil.AssertEqual(t, 10002, as.GetCurrentValue())
	as.Add(-1)
	testutil.AssertEqual(t, 10001, as.GetCurrentValue())
}

// Tests for NewDog

func TestNewDog_MakeNoise_ShouldMakeABarkerNoise(t *testing.T) {
	dog := NewDog()

	testutil.AssertEqual(t, "BARK BARK!!!!", dog.MakeNoise())
}

func TestNewDog_RollOverWhileGood_ShouldRollOver(t *testing.T) {
	dog := NewDog()
	dog.SetIsGoodDog(true)

	testutil.AssertEqual(t, true, dog.RollOver())
}

func TestNewDog_RollOverWhileBad_ShouldNotRollOver(t *testing.T) {
	dog := NewDog()
	dog.SetIsGoodDog(false)

	testutil.AssertEqual(t, false, dog.RollOver())
}

func TestNewDog_ShouldContainBarker(t *testing.T) {
	dog := NewDog()

	v := reflect.ValueOf(dog)
	i := reflect.Indirect(v)
	foundBarker := false
	for itr := 0; itr < i.NumField(); itr++ {
		if "structs.Barker" == i.Field(0).Type().String() {
			foundBarker = true
		}
	}
	testutil.AssertEqualMsg(t, true, foundBarker, "You need to embed the Barker struct in your struct (if you embedded "+
		"Barker and this test isn't passing try making the embedded Barker the first field in your Dog impl struct)")
}
