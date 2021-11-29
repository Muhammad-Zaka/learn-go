package main

import (
	"fmt"

	"github.com/Muhammad-Zaka/learn-go/handlers"
)

func main() {

	const PI = 3.14
	radius := 5
	circumference := 2 * PI * float64(radius)
	fmt.Printf("For a radius of %v. The circumference is %v", radius, circumference)
	fmt.Println(handlers.Handle)
}
