package main

import "fmt"

func main() {
	age := 32
	myage := &age
	doubledAge := double(myage)
	fmt.Println(doubledAge)

}

func double(number *int) int {
	return *number * 2
}
