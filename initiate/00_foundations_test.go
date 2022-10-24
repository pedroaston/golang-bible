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

func TestBasicNumericVariables(t *testing.T) {

	// Defining a variable without initializing it at the
	// int type can be of 8 bits (-128 to 127), 16, 32, 64
	// uint type can have the same bit size but only has positive values
	// float type can only be 32 or 64 bits
	var anInt int8
	var anUnsignedInt uint8
	var aFloat float32

	// Define variable by initializing it with a value
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
}
