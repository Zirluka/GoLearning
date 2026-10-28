package main

import "fmt"

func main() {
	// short create -> varName := ...
	shortCreate := 111
	// change var -> varName = ...
	shortCreate = 50
	// long create -> var varName type = ...
	var longCreate float64 = 222.5
	// or empty var -> var varName type
	var changeAfter string
	changeAfter = "changed"

	fmt.Println(shortCreate)
	fmt.Println(longCreate)
	fmt.Println(changeAfter)
}