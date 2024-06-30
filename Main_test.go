package main

import (
	"testing"
)

func TestStartsWithEmptyStack(t *testing.T) {
	_, err := StackMachine("")

	if err == nil {
		t.Error("expected error due to no results")
	}
}

// Write your own TDD tests here as you develop
func TestSingleIntegerInCommand(t *testing.T) {
	result, err := StackMachine("13")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if result != 13 {
		t.Error("Expected 13, got", result)
	}
}

func TestMultipleIntigersInCommand(t *testing.T) {
	result, err := StackMachine("13 7 9")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if result != 9 {
		t.Error("Expected 9, got ", result)
	}
}
func TestInputtingNegativeNumberReturnsError(t *testing.T) {
	_, err := StackMachine("-1")
	if err == nil {
		t.Error("Expected an error, did not receive one.")
	}
}

func TestInputtingNumberAbove50000ReturnsError(t *testing.T) {
	_, err := StackMachine("50001 20 3")
	if err == nil {
		t.Error("Expected an error, did not receive one.")
	}
}

func TestEmptyStringReturnsError(t *testing.T) {
	_, err := StackMachine(" ")
	if err == nil {
		t.Error("Expected error from inputting an empty string.")
	}
}

func TestAddition(t *testing.T) {
	result, err := StackMachine("13 7 +")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if result != 20 {
		t.Error("Expected 20, got ", result)
	}
}

func TestOneNumberAndAdditionSignReturnsError(t *testing.T) {
	_, err := StackMachine("1 +")
	if err == nil {
		t.Error("Expected an error.")
	}
}
func TestAdditionResultOver50kReturnsErr(t *testing.T) {
	_, err := StackMachine("50000 1 +")
	if err == nil {
		t.Error("Expected error, did not get one.")
	}
}

func TestAddingThe2ndAnd3rdNumberWorks(t *testing.T) {
	result, err := StackMachine("2 7 3 +")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 10 {
		t.Error("Expected 10, got: ", result)
	}
}
func TestAddingMoreThanOnce(t *testing.T) {
	result, err := StackMachine("4 4 + 4 +")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 12 {
		t.Error("Expected 12, got: ", result)
	}
}

func TestSubtraction(t *testing.T) {
	result, err := StackMachine("7 13 -")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if result != 6 {
		t.Error("Expected 6, got ", result)
	}
}

func TestOneNumberAndSubtractionSignReturnsError(t *testing.T) {
	_, err := StackMachine("1 -")
	if err == nil {
		t.Error("Expected an error.")
	}
}
func TestNegativeSubtractionResultReturnsError(t *testing.T) {
	_, err := StackMachine("7 2 -")
	if err == nil {
		t.Error("Expected an error, did not get one.")
	}
}

func TestSubtractingThe2ndAnd3rdNumberWorks(t *testing.T) {
	result, err := StackMachine("2 3 7 -")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 4 {
		t.Error("Expected 10, got: ", result)
	}
}
func TestSubtractingMoreThanOnce(t *testing.T) {
	result, err := StackMachine("4 6 - 3 -")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 1 {
		t.Error("Expected 4, got: ", result)
	}
}

func TestMultiplication(t *testing.T) {
	result, err := StackMachine("2 3 *")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if result != 6 {
		t.Error("Expected 6, got ", result)
	}
}
func TestOneNumberAndMultiplicationSignReturnsError(t *testing.T) {
	_, err := StackMachine("1 *")
	if err == nil {
		t.Error("Expected an error.")
	}
}

func TestMultiplicationResultOver50kReturnsError(t *testing.T) {
	_, err := StackMachine("16667 3 *") //gives the result of 50,001 exactly
	if err == nil {
		t.Error("Expected error, did not get one.")
	}
}

func TestMultiplying2ndAnd3rdNumbersWork(t *testing.T) {
	result, err := StackMachine("7 8 2 *")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 16 {
		t.Error("Expected 16, got ", result)
	}
}
func TestMultiplyingMoreThanOnce(t *testing.T) {
	result, err := StackMachine("4 4 * 4 *")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 64 {
		t.Error("Expected 64, got: ", result)
	}
}
func TestMultiplyingAndAdding(t *testing.T) {
	result, err := StackMachine("8 2 + 2 *")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 20 {
		t.Error("Expected 20, got: ", result)
	}
}
func TestAddingAndSubtracting(t *testing.T) {
	result, err := StackMachine("8 2 + 17 -")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 7 {
		t.Error("Expected 7, got: ", result)
	}
}

func TestSubtractingAndMultiplying(t *testing.T) {
	result, err := StackMachine("8 2 * 28 -")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 12 {
		t.Error("Expected 12, got: ", result)
	}
}
func TestAddingSubtractingANDMultiplying(t *testing.T) {
	result, err := StackMachine("8 2 + 18 - 8 *")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 64 {
		t.Error("Expected 64, got: ", result)
	}
}
func TestDuplication(t *testing.T) {
	result, err := StackMachine("3 DUP")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 3 {
		t.Error("Expected 3, got: ", result)
	}
}
func TestDuplicationWith2NumbersInCommand(t *testing.T) {
	result, err := StackMachine("4 3 DUP")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 3 {
		t.Error("Expected 3, got: ", result)
	}
}
func TestCanDuplicateNothing(t *testing.T) {
	result, err := StackMachine("DUP 3")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 3 {
		t.Error("Expected 3, got: ", result)
	}
}

func TestCanDuplicateSums(t *testing.T) {
	result, err := StackMachine("3 4 + DUP")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 7 {
		t.Error("Expected 7, got: ", result)
	}
}

func TestCannotDuplicateNumberOver50k(t *testing.T) {
	_, err := StackMachine("50001 DUP")
	if err == nil {
		t.Error("Expected an error, got none.")
	}
}

func TestPop(t *testing.T) {
	result, err := StackMachine("3 4 POP")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 3 {
		t.Error("Expected 3, got: ", result)
	}
}

func TestPuttingPopBeforeNothingDoesntCausePanicAndReturnsErr(t *testing.T) {
	_, err := StackMachine("POP 8")
	if err == nil {
		t.Error("Expected error, got none.")
	}
}

func TestSum(t *testing.T) {
	result, err := StackMachine("1 2 3 4 5 SUM")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 15 {
		t.Error("Expected 15, got: ", result)
	}
}
func TestCantMakeSumOfNothing(t *testing.T) {
	_, err := StackMachine("SUM 3")
	if err == nil {
		t.Error("Expected an error, got none.")
	}
}

func TestSumOver50kReturnsErr(t *testing.T) {
	_, err := StackMachine("49998 1 2 SUM")
	if err == nil {
		t.Error("Expected an error, got none.")
	}
}

func TestClear(t *testing.T) {
	result, err := StackMachine("1 2 SUM CLEAR 2 2 SUM")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if result != 4 {
		t.Error("Expected 4, got: ", result)
	}
}

/*
All these tests must pass for completion
*/
/*func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		name        string
		commands    string
		expected    int
		expectedErr error
	}{
		{name: "empty error", commands: "", expected: 0, expectedErr: errors.New("")},
		{name: "add overflow", commands: "50000 DUP +", expected: 0, expectedErr: errors.New("")},
		{name: "too few add", commands: "99 +", expected: 0, expectedErr: errors.New("")},
		{name: "too few minus", commands: "99 -", expected: 0, expectedErr: errors.New("")},
		{name: "too few multiply", commands: "99 *", expected: 0, expectedErr: errors.New("")},
		{name: "empty stack", commands: "99 CLEAR", expected: 0, expectedErr: errors.New("")},
		{name: "sum single value", commands: "99 SUM", expected: 99, expectedErr: nil},
		{name: "sum empty", commands: "SUM", expected: 0, expectedErr: errors.New("")},
		{name: "normal +*", commands: "5 6 + 2 *", expected: 22, expectedErr: nil},
		{name: "clear too few", commands: "1 2 3 4 + CLEAR 12 +", expected: 0, expectedErr: errors.New("")},
		{name: "normal after clear", commands: "1 CLEAR 2 3 +", expected: 5, expectedErr: nil},
		{name: "single integer", commands: "9876", expected: 9876, expectedErr: nil},
		{name: "invalid command", commands: "DOGBANANA", expected: 0, expectedErr: errors.New("")},
		{name: "normal +-*", commands: "5 9 DUP + + 43 - 3 *", expected: 60, expectedErr: nil},
		{name: "minus", commands: "2 5 -", expected: 3, expectedErr: nil},
		{name: "underflow minus", commands: "5 2 -", expected: 0, expectedErr: errors.New("")},
		{name: "at overflow limit", commands: "25000 DUP +", expected: 50000, expectedErr: nil},
		{name: "at overflow limit single value", commands: "50000 0 +", expected: 50000, expectedErr: nil},
		{name: "overflow plus", commands: "50000 1 +", expected: 0, expectedErr: errors.New("")},
		{name: "overflow single value", commands: "50001", expected: 0, expectedErr: errors.New("")},
		{name: "times zero at overflow limit", commands: "50000 0 *", expected: 0, expectedErr: nil},
		{name: "too few at first", commands: "1 2 3 4 5 + + + + * 999", expected: 0, expectedErr: errors.New("")},
		{name: "normal simple", commands: "1 2 - 99 +", expected: 100, expectedErr: nil},
		{name: "at overflow minus to zero", commands: "50000 50000 -", expected: 0, expectedErr: nil},
		{name: "clear empties stack", commands: "CLEAR", expected: 0, expectedErr: errors.New("")},
		{name: "normal sum", commands: "3 4 3 5 5 1 1 1 SUM", expected: 23, expectedErr: nil},
		{name: "sum after clear stack", commands: "3 4 3 5 CLEAR 5 1 1 1 SUM", expected: 8, expectedErr: nil},
		{name: "sum then too few", commands: "3 4 3 5 5 1 1 1 SUM -", expected: 0, expectedErr: errors.New("")},
		{name: "fibonacci", commands: "1 2 3 4 5 * * * *", expected: 120, expectedErr: nil},
	}

	for _, test := range tests {

		got, err := StackMachine(test.commands)

		if test.expectedErr != nil {
			if err == nil {
				t.Errorf("%s (%s) Expected error, but received nil", test.name, test.commands)
			} else if err.Error() != test.expectedErr.Error() {
				t.Errorf("%s (%s) got error %v, want %v", test.name, test.commands, err, test.expectedErr)
			}
		} else if got != test.expected {
			t.Errorf("%s (%s) got %v, want %v", test.name, test.commands, got, test.expected)
		}
	}
}*/
