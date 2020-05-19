package main

import (
	"math/rand"
	"time"
)

// import "fmt"

// var name string
// var age int
var height float64

func main() {
	seedvalue := time.Now().Unix()
	rand.Seed(seedvalue)
	// nameage("Kavish", 40)
	// var a string
	var a int
	a = rannum()
	print(a,"\n")
}
