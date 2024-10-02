package main

import (
	"fmt"
	"math/rand"
)

func main() {
	println("Hello World!")
	print("Hi folks, lets learn Go")
	fmt.Println("Hello Go minds, welcome to go using fmt pacakge")
	println("Randon number:", rand.Intn(100))
	//os.Exit(1)
}

// main.main

// 1. create a package --> library
// 2. create a binary/executable

// compile-time vs runtime
// statically typed programming language

/*
 - run go application
 - 1. it is compiled --> compiler
 - 2. It creates assembly code (plan9 assembly)
 - 3. assembler -> it concerts the code into object code
 - 4. linker -> links all required things and creates a binary/executable
 - 5. the created executable is executed

 - build go application
 - 1. it is compiled --> compiler
 - 2. It creates assembly code (plan9 assembly)
 - 3. assembler -> it concerts the code into object code
 - 4. linker -> links all required things and creates a binary/executable
*/
