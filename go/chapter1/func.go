package main

import (
	"bufio"
	"fmt"
	"os"
)

func nameage(name string, age int) {
	// name = "Kavish"
	// age = 11
	height = 147.123
	print("Hi ", name, ". Your age is ", age, newline())
	print("Height : ", height, newline())
	fmt.Printf("My Heigt %5.2fm\n", height)
	fmt.Printf("My age [%03d] \n", age)
	print("Height : ", height)

}
func newline() string {
	return "\n"
}

func inputname() (name string, age int) {
	// var name string
	// var age int
	// fmt.Printf("My age [%03d] \n", age)
	fmt.Printf("Enter Name : ")
	in := bufio.NewReader(os.Stdin)
	name, _ = in.ReadString('\n')
	fmt.Printf("Enter Age : ")
	fmt.Scan(&age)
	print("Hi ", name,". Your age is ", age, newline())

	return name, age
}