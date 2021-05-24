package main

import (
	"fmt"
	"math"
)

func main() {
	var age int

	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after initialisation:", age)

	var height int = 181
	fmt.Println("My height is:", height)

	var uid = 1811231
	fmt.Println("My uid is:", uid)
	fmt.Printf("%T\n", uid)

	var a, b = 30, "Vova"
	fmt.Printf("type a = %T\n", a)
	fmt.Printf("type b = %T\n", b)

	count := 10
	fmt.Println("count = ", count)
	count++
	fmt.Println("count = ", count)

	width, lenght := 20.5, 30.2
	fmt.Printf("Min dimensional of rectangle is: %.3f\n", math.Min(width, lenght))

}
