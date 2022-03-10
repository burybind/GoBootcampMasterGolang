package main

import (
	"fmt"
	"os"
)

func main() {
	age()
	name()
	claims()
	temp()
	quotes()
	type1()
	type2()
	type3()
	type4()
	fullname()
}

// ---------------------------------------------------------
// EXERCISE: Print Your Age
//
//  Print your age using Printf
//
// EXPECTED OUTPUT
//  I'm 30 years old.
//
// NOTE
//  You should change 30 to your age, of course.
// ---------------------------------------------------------

func age() {
	age := 34
	fmt.Printf("My age: %d\n", age)
}

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func name() {
	// BONUS: Use a variable for the format specifier
	str := "My name is %s and my lastname is %s\n"
	fname := "Inanc"
	lname := "Gumus"
	fmt.Printf(str, fname, lname)
}

// ---------------------------------------------------------
// EXERCISE: False Claims
//
//  Use printf to print the expected output using a variable.
//
// EXPECTED OUTPUT
//  These are false claims.
// ---------------------------------------------------------

func claims() {
	// UNCOMMENT THE FOLLOWING CODE
	// AND DO NOT CHANGE IT AFTERWARDS
	tf := false

	// TYPE YOUR CODE HERE
	fmt.Printf("These are %t claims\n", tf)
}

// ---------------------------------------------------------
// EXERCISE: Print the Temperature
//
//  Print the current temperature in your area using Printf
//
// NOTE
//  Do not use %v verb
//  Output "shouldn't" be like 29.500000
//
// EXPECTED OUTPUT
//  Temperature is 29.5 degrees.
// ---------------------------------------------------------

func temp() {
	temp := 29.5
	fmt.Printf("Temperature is %.1f degrees.\n", temp)
}

// ---------------------------------------------------------
// EXERCISE: Double Quotes
//
//  Print "hello world" with double-quotes using Printf
//  (As you see here)
//
// NOTE
//  Output "shouldn't" be just: hello world
//
// EXPECTED OUTPUT
//  "hello world"
// ---------------------------------------------------------

func quotes() {
	str := "hello world"
	fmt.Printf("%q\n", str)
}

// ---------------------------------------------------------
// EXERCISE: Print the Type
//
//  Print the type and value of 3 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3 is int
// ---------------------------------------------------------

func type1() {
	v := 3
	fmt.Printf("Type of %d is %T\n", v, v)
}

// ---------------------------------------------------------
// EXERCISE: Print the Type #2
//
//  Print the type and value of 3.14 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3.14 is float64
// ---------------------------------------------------------

func type2() {
	v := 3.14
	fmt.Printf("Type of %.2f is %T\n", v, v)
}

// ---------------------------------------------------------
// EXERCISE: Print the Type #3
//
//  Print the type and value of "hello" using fmt.Printf
//
// EXPECTED OUTPUT
// 	Type of hello is string
// ---------------------------------------------------------

func type3() {
	v := 3
	fmt.Printf("Type of %d is %T\n", v, v)
}

// ---------------------------------------------------------
// EXERCISE: Print the Type #4
//  Print the type and value of true using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of true is bool
// ---------------------------------------------------------

func type4() {
	v := true
	fmt.Printf("Type of %t is %T\n", v, v)
}

// ---------------------------------------------------------
// EXERCISE: Print Your Fullname
//
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
//
// EXAMPLE INPUT
//  Inanc Gumus
//
// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.
// ---------------------------------------------------------

func fullname() {
	// BONUS: Use a variable for the format specifier
	fname := os.Args[1]
	lname := os.Args[2]
	str := "Your name is %s and your lastname is %s"
	fmt.Printf(str, fname, lname)
}
