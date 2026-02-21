package main

import "fmt"

func main() {
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}

/*

CHAPTER 1
:= is called the walrus operator
Go can infer the type using := is <type> because of the first value (type inference)

Go Features
Fast and Lighweight, Easily Concurrent, Easy and Simple, Compiled, Statically Typed, Garbage Collected

Go Compiler - take go code and produce machine code like .exe

package main - Go compiler know that we want this code to compile and run as a standalone program
import "fmt" imports the formatting package from the standard library
func main() defines the main function, entry point of Go

2 kinds of errors
compilation errors - better since we can catch it before going into prod
runtime errors - worse since they are unexpected

type sizes
int int8 int16 int32 int64
uint ... uintptr
float32 float64
complex64 complex128

we'll use more spcific types when we are concerned about performance and memory usage

go is statically typed
*/