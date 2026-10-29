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
}
