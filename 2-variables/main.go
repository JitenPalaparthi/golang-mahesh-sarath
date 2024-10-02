package main

import "fmt"

func main() {
	var age1 uint8 = 39 // 1

	var age2 uint8 // automatically 0 is inferred

	var age3 = 39 // bad approach still ppl use it.

	age4 := 39 // shorthand declaration

	var amount1 = 23.45

	amount2 := 234.234

	// var ok1 bool = false
	// var ok2 = false
	// ok3 := true

	println("age1:", age1, "age2:", age2, "age3:", age3, "age4:", age4)
	fmt.Printf("\nValue of age3:%v type of age3:%T", age3, age3)
	fmt.Printf("\nValue of age4:%v type of age4:%T", age4, age4)
	fmt.Printf("\nValue of amount1:%v type of amount1:%T", amount1, amount1)
	fmt.Printf("\nValue of amount2:%v type of amount2:%T", amount2, amount2)
}

// primitive/ defined types
// numbers -> int,uint,int8,int16,in32,int64,uing8,uint16,uint32,uint64,float32,float64
// bool
// string

// 12312.1232 + 12312.12312

// uint8 + uint8
// 255 + 255   -> uint16

// process is created
// Memory, thread

// Code Segment --> const variable
// Data Segment --> Global variables
// Stack Memory -->
// Heap Memory  -->
