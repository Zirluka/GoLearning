package main

import "fmt"

func main() {
	// defer - это отложенная функция, как в <script defer> в html
	defer func() {
		fmt.Println("Defer func")
	}() // -> () в конце это моментальный вызов фунцкии

	fmt.Println("Hello1")
	fmt.Println("Hello2")
	fmt.Println("Hello3")
	foo()
}

func foo() {
	// Так как defer использует структуру данных stack, defer'ы будут идти снизу вверх, то есть 3 -> 2 -> 1
	defer func() {
		fmt.Println("Defer 1")
	}()
	defer func() {
		fmt.Println("Defer 2")
	}() 
	defer func() {
		fmt.Println("Defer 3")
	}() 

	fmt.Println("foo")
}