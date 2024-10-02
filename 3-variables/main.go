package main

import "fmt"

var (
	r1 float32 = 25.25    // data segment
	r2 float32 = square() // data segment
)

func main() {

	const PI float32 = 3.14 // code segment

	const PI1 = 3.14 // shorthand declaration ; code segment

	//var num uint8 = 2
	const PI1OF2 = PI * 2    // evaluated by the compiler and assigns a value to this const
	const PI1OF21 = PI + PI1 // evaluated by the compiler and assigns a value to this const

	var r float32 = 25.35 // stack allocated
	fmt.Printf("\nValue of PI1: %v type of PI1: %T", PI1, PI1)

	//r3 float32 = 25.34 // data segment

	println("\n2pir=1=", 2*PI*r)

	println("\n 2pir-2=", 2*PI1*r1)

	//r2 = 12.34
	r1 = 12.343 * r2
	r = 25.45 * 12.0 * r1
	//v1 := square()
	//	println("v1 square:", v1)
	{
		r2 := 123.213
		//println("r2:", r2)
		{
			r2 := 1.12
			fmt.Printf("r2: %.3f\n", r2)
		}
		fmt.Printf("r2: %.3f\n", r2)
	}
	fmt.Printf("r2: %.3f", r2)
}

func square() float32 {
	return r1 * r1
	//return r * r
}

// constants are stored in code segment
//
