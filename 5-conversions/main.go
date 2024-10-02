package main

import "fmt"

func main() {

	var (
		num1 uint8
		num2 uint64
	)

	num1 = 200

	// num2 = num1 // this is not okay in Go, implicit conversion is not allowed

	num2 = uint64(num1) // type casting in Go

	println("num1:", num1, "num2:", num2) // built in function

	var (
		float1 float64 = 123.123123
		float2 float32 = float32(float1) // type caste
	)

	fmt.Println("float2:", float2) // function from a fmt package

	ok := true // shorthand declaration
	println("ok:", ok)
	// var num3 uint8 = uint8(ok) // not ; though both carry 1 byte , casting is not allowed

	var (
		num4   uint8                         // 0
		num5   uint64  = 2000034234234324234 // 8 bytes
		float3 float32 = 2312.232
		num6   uint32  = uint32(float3)
	)

	num4 = uint8(num5)

	println("num4:", num4)
	println("float3 to num5", num6)

}

// int --> 8 bytes (on 64 bit machines)
// float32 --> 4 bytes
// byte --> 1 byte

// int8  -- 1 byte can easily be converted to int (8 bytes)
