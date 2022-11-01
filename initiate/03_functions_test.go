package initiate

import (
	"fmt"
	"testing"
)

func SimpleFunction() {
	fmt.Printf("%d + %d = %d\n", 3, 2, 3+2)
}

func TestSimpleFunction(t *testing.T) {

	// Executes one simple functions
	SimpleFunction()
}

func OneArgOneReturnedValue(x int) int {
	return x * 2
}

func FourArgsTwoReturnedValues(x, y int, name string, word string) (int, string) {

	return x + y, name + word
}

func TestMoreComposedFunctions(t *testing.T) {

	res := OneArgOneReturnedValue(2)
	res1, res2 := FourArgsTwoReturnedValues(2, 3, "hello", "world")
	fmt.Printf("%d | %d %s\n", res, res1, res2)
}
