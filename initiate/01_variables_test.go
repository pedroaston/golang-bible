package initiate

import (
	"fmt"
	"testing"
)

/*
	This file has the foundation for the next files of the initiate package

	Helper: To run a specific testcase you can use vscode's run test
	or debug test if you want to execute via terminal command use
	the following instruction on a terminal console at this directory

	go test -run VariableBasics

	Note: The -run flag executes every testcase that follows a certain regex
*/

// Global Variable
// can be changed
var globalInt int = 2

// Constant
// CANNOT be changed
const pi = 3.1415

func TestBasicNumericVariables(t *testing.T) {

	// Defining a variable without initializing it at the
	// int type can be of 8 bits (-128 to 127), 16, 32, 64
	// uint type can have the same bit size but only has positive values
	// float type can only be 32 or 64 bits
	var anInt int8
	var anUnsignedInt uint8
	var aFloat float32

	// Define variable by initializing it with a value
	// This is called "Typed Inference"
	anotherNumber := 7

	// Attribute value to the variable
	anInt = 1
	anUnsignedInt = 2
	aFloat = 1.3
	anotherNumber += 1

	// Increment a number by one
	anInt++
	aFloat++

	// Just printing the variables because they need to be use at
	// least once or else the code will fail to build
	fmt.Println(anInt)
	fmt.Println(anotherNumber)
	fmt.Println(anUnsignedInt)
	fmt.Println(aFloat)
	fmt.Println(globalInt)
	fmt.Println(pi)
}

// TestDefaultValues shows some variables
// default values upon assigning type
func TestDefaultValues(t *testing.T) {

	var intNumber int
	var boolean bool
	var floatNumber float32
	var oneStr string

	fmt.Println(intNumber)
	fmt.Println(boolean)
	fmt.Println(floatNumber)
	// The string is "" so it is shown an empty line
	fmt.Println(oneStr)
}

// TestTypeConversion
func TestTypeConvertion(t *testing.T) {

	var anInt int
	var aFloat float32

	anInt = 2
	aFloat = float32(anInt)
	fmt.Println(aFloat)

	// anInt = int(aStr)
	// The above line isn't executable
	// there are additional packages that do this conversion
}
