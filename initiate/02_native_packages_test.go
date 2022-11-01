package initiate

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

/*
	This file has some essential functions from golang native
	packages, that will be useful for your golang journey
*/

// TestGolangFMT just showcases some of the most relevant
// code of the fmt package and some string encoding
func TestGolangFMT(t *testing.T) {

	// Printing test on terminal
	fmt.Println("This prints a line")
	fmt.Printf("This prints the integer %d \n", 1)
	fmt.Printf("This prints the float %f \n", 3.14)
	fmt.Printf("Is this relevant: %v \n", false)
	fmt.Printf("Data-type %T | Data-value %v\n", 3.14, 3.14)

	// Generating Strings
	newstr := fmt.Sprintf("This string, plus this one%s plus this number %d!", ",", 666)

	// Concatenating Strings
	fmt.Println("Guess what, + concatenates strings :o " + newstr)
}

// TestGolangTime showcases some of the uses of the time package like
// betting current time, operations with several time/Duration instances,
// how to make a thread inactive for a certain amount of time, ...
func TestGolangTime(t *testing.T) {

	// Returns current time
	aTimeInstance := time.Now()
	fmt.Println(aTimeInstance)

	// Adding several time values
	fmt.Println(aTimeInstance.Add(2*time.Second + 3*time.Minute + 2*time.Hour))

	// Make thread sleep for a certain time
	aTimeStamp := time.Now()
	time.Sleep(500 * time.Millisecond)
	sleepTime := time.Since(aTimeStamp)
	fmt.Println(sleepTime)
}

// TestGolangMath explores some of the interesting function of
// the package math, such as the random number generation
func TestGolangMath(t *testing.T) {

	// Generates a random number from a particular
	// interval 0 to n (=100 in this case)
	fmt.Printf("Random number between 0-100: %d\n", rand.Intn(100))

	// The above number is always the same because it is used
	// the same initializing generator value. To have a constantly
	// different number use a different seed value before generating
	// you first random number. Time is a good seed candidate.
	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Printf("Almost allways different number: %d\n", rand.Intn(100))

	// This package also has several mathematical functions
	fmt.Println(math.Cos(math.Pi))
}
