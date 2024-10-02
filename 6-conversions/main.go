package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	type integer = int // creating alias
	type char = rune

	var (
		//byte1 byte // alias of uint8
		char1 rune // alias of int32
		char2 char // alias of rune which is an alias of int32
		char3 char
		char4 int32
		char5 rune = 'A'
	)

	// 0 - 65,XXX
	// 0-255

	// ascii vs utf8

	println("你好世界")

	char1 = '你' // utf8 each char holds 4 byte

	char2 = '好'

	char3 = char2 // not converting one type to the other type , rather just using different name to the same type

	char4 = char1

	println("char1:", char1)
	println("char2:", char2)
	println("char3:", char3)
	println("char4:", char4)
	println("char5:", char5)

	println(string(char1))

	var (
		val1 string = "12321"
		num1 int
	)

	num1, _ = strconv.Atoi(val1) // will explain multiple return types later
	println("num1:", num1)

	val1 = "123A0"
	num1, err := strconv.Atoi(val1) // will explain multiple return types later
	fmt.Println("num1:", num1, "Error?", err)

	num1 = 1231231

	// convert number to string

	str1 := strconv.Itoa(num1) // number to string

	fmt.Println("str1:", str1, "type:", reflect.TypeOf(str1))

	fmt.Printf("str1: %s type:%T\n", str1, str1) // both are same %T and reflect.TypeOf

	ok := true

	float1 := 12312.1231231

	str2 := fmt.Sprint("num1: ", num1, "float1: ", float1, "ok: ", ok) // kind of a formatter
	fmt.Printf("str2: %s type:%T\n", str2, str2)                       // both are same %T and reflect.TypeOf

}

// nil
