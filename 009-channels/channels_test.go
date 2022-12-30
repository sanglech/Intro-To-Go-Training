package channels

import (
	"introduction-to-go-training/testutil"
	"testing"
	"time"
)

// Tests for SendRange

func TestSendRange_StartAfterEnd_SendsNothing(t *testing.T) {
	output := make(chan int)

	go SendRange(10, 9, output)

	select {
	case _, ok := <-output:
		testutil.AssertEqual(t, false, ok)
	case <-time.After(1 * time.Second):
		t.Fatal("Timed out waiting for channel to be closed.")
		return
	}
}

func TestSendRange_StartEqualsEnd_SendsOneNumber(t *testing.T) {
	output := make(chan int)

	go SendRange(10, 10, output)

	var results []int
	for {
		select {
		case first, ok := <-output:
			if !ok {
				goto ASSERTIONS
			}
			results = append(results, first)
		case <-time.After(1 * time.Second):
			t.Fatal("Timed out waiting for channel to be closed.")
			return
		}
	}

ASSERTIONS:
	expected := []int{10}
	testutil.AssertEqualMsgf(t, len(expected), len(results), "Expected length %v but got length %v")
	for itr, val := range results {
		testutil.AssertEqualMsgf(t, expected[itr], val, "Expected %v but got %v (slice index %v)", itr)
	}
}

func TestSendRange_StartAfterEnd_SendsIntRange(t *testing.T) {
	output := make(chan int)

	go SendRange(-5, 3, output)

	var results []int
	for {
		select {
		case first, ok := <-output:
			if !ok {
				goto ASSERTIONS
			}
			results = append(results, first)
		case <-time.After(1 * time.Second):
			t.Fatal("Timed out waiting for channel to be closed.")
			return
		}
	}

ASSERTIONS:
	expected := []int{-5, -4, -3, -2, -1, 0, 1, 2, 3}
	testutil.AssertEqualMsgf(t, len(expected), len(results), "Expected length %v but got length %v")
	for itr, val := range results {
		testutil.AssertEqualMsgf(t, expected[itr], val, "Expected %v but got %v (slice index %v)", itr)
	}
}

// Tests for BuildString

func TestBuildString_SendNothing_ReturnsEmptyString(t *testing.T) {
	input := make(chan string)
	output := make(chan string)

	go BuildString(input, output)

	close(input)

	select {
	case res := <-output:
		testutil.AssertEqual(t, "", res)
	case <-time.After(1 * time.Second):
		t.Fatal("Timed out waiting for channel to be closed.")
		return
	}
}

func TestBuildString_SendOneString_ReturnsTheSentString(t *testing.T) {
	input := make(chan string, 1)
	output := make(chan string)

	go BuildString(input, output)

	input <- "Hello"
	close(input)

	select {
	case res := <-output:
		testutil.AssertEqual(t, "Hello", res)
	case <-time.After(1 * time.Second):
		t.Fatal("Timed out waiting for channel to be closed.")
		return
	}
}

func TestBuildString_SendSeveralStrings_ReturnsConcatenation(t *testing.T) {
	input := make(chan string, 5)
	input <- "Hello"
	input <- " how's it going"
	input <- " my duuuuude"
	input <- "?"
	close(input)
	output := make(chan string)

	go BuildString(input, output)

	select {
	case res := <-output:
		testutil.AssertEqual(t, "Hello how's it going my duuuuude?", res)
	case <-time.After(1 * time.Second):
		t.Fatal("Timed out waiting for channel to be closed.")
		return
	}
}

// Test ConvertToLowercase

func TestConvertToLowercase_SendNothingAndCloseChannel_ReceiveNothing(t *testing.T) {
	input := make(chan string)
	close(input)
	interrupt := make(chan struct{}, 1)
	output := make(chan string, 1)

	go ConvertToLowercase(input, output, interrupt)

	select {
	case _, ok := <-output:
		testutil.AssertEqual(t, false, ok)
	case <-time.After(1 * time.Second):
		t.Fatal("Timed out waiting for channel to be closed.")
		return
	}
}

func TestConvertToLowercase_SendNothingAndInterrupt_ReceiveNothing(t *testing.T) {
	input := make(chan string)
	interrupt := make(chan struct{}, 1)
	interrupt <- struct{}{}
	output := make(chan string, 1)

	go ConvertToLowercase(input, output, interrupt)

	select {
	case _, ok := <-output:
		testutil.AssertEqual(t, false, ok)
	case <-time.After(1 * time.Second):
		t.Fatal("Timed out waiting for channel to be closed.")
		return
	}
}

func TestConvertToLowercase_SendStringsAndClose_ReceiveLowercaseStrings(t *testing.T) {
	input := make(chan string, 3)
	input <- "YO"
	input <- "BIG CAPS"
	close(input)
	interrupt := make(chan struct{}, 1)
	output := make(chan string, 1)

	go ConvertToLowercase(input, output, interrupt)

	var results []string
	for {
		select {
		case first, ok := <-output:
			if !ok {
				goto ASSERTIONS
			}
			results = append(results, first)
		case <-time.After(1 * time.Second):
			t.Fatal("Timed out waiting for channel to be closed.")
			return
		}
	}

ASSERTIONS:
	expected := []string{"yo", "big caps"}
	testutil.AssertEqualMsgf(t, len(expected), len(results), "Expected length %v but got length %v")
	for itr, val := range results {
		testutil.AssertEqualMsgf(t, expected[itr], val, "Expected %v but got %v (slice index %v)", itr)
	}
}

func TestConvertToLowercase_SendStringsAndInterrupt_ReceiveLowercaseStringsFromBeforeInterrupt(t *testing.T) {
	input := make(chan string, 2)
	interrupt := make(chan struct{}, 1)
	output := make(chan string, 5)

	go ConvertToLowercase(input, output, interrupt)

	input <- "YO"
	input <- "BIG CAPS"

	// block here to switch control to the other goroutine
	firstRes := ""
	select {
	case val := <-output:
		firstRes = val
	case <-time.After(1 * time.Second):
		t.Fatal("Timed out waiting for channel to send for the first time.")
		return
	}

	results := []string{firstRes}
	interrupt <- struct{}{}

	// block here to switch control to the other goroutine
	secondRes := ""
	select {
	case val := <-output:
		secondRes = val
	case <-time.After(1 * time.Second):
		t.Fatal("Timed out waiting for channel to send for the second time.")
		return
	}
	results = append(results, secondRes)

	input <- "BIG CAPS2"

	// verify that the output channel is closed rather than "big caps2" being processed
	select {
	case _, ok := <-output:
		testutil.AssertEqual(t, false, ok) // should be closed
	case <-time.After(1 * time.Second):
		t.Fatal("Timed out waiting for channel to close.")
		return
	}

	expected := []string{"yo", "big caps"}
	testutil.AssertEqualMsgf(t, len(expected), len(results), "Expected length %v but got length %v")
	for itr, val := range results {
		testutil.AssertEqualMsgf(t, expected[itr], val, "Expected %v but got %v (slice index %v)", itr)
	}
}
