package main

import "math/rand"

func main() {
	a := 5
	b := 5
	sum := add(a, b)
	println(sum)
	println(print(1))
	c, d := generateRandomNumbers()
	println(c, d)
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func print(num1 int) int {
	return num1
}

func generateRandomNumbers() (rand1 int, rand2 int) {
	rand1 = rand.Intn(10000)
	rand2 = rand.Intn(10000)
	return
}
