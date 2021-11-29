package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	// Output information
	fmt.Println("BMI Calculator")
	fmt.Println("--------------------")
	// Prompt for user input (weight + height)
	fmt.Println("Please enter your weight (kg): ")
	inputWeight, _ := reader.ReadString('\n')
	fmt.Println("Please enter your height (m): ")
	inputHeight, _ := reader.ReadString('\n')

	// Save that user input in variables
	inputWeight = strings.TrimSuffix(inputWeight, "\r\n")
	inputHeight = strings.TrimSuffix(inputHeight, "\r\n")
	weight, _ := strconv.ParseFloat(inputWeight, 64)
	height, _ := strconv.ParseFloat(inputHeight, 64)
	fmt.Println(inputHeight)
	fmt.Println(inputWeight)
	fmt.Println(height)
	fmt.Println(weight)
	// Calculate the BMI (weight /(height*height))
	bmi := weight / (height * height)
	// Output the calcluated BMI
	fmt.Printf("Your BMI: %.2f", bmi)
}
