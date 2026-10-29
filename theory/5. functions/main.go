package main

import "fmt"

func main() {
	printSquareOfInt(3)
}

func square(x int) int {
	return x * x
}

func printSquareOfInt(x int) {
	fmt.Println("Квадрат", x, "равен", square(x))
}