package main

import "fmt"

func main() {

	PI := 3.14
	radius := 5
	circumference := 2 * PI * float64(radius)

	fmt.Printf("For a radius of %v. The circumference is %v", radius, circumference)
}