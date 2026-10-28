package main

import "fmt"

func main() {
	score := 45

	// if-else construction
	if score >= 15 {
		fmt.Println("Ты мега-красавчик!")
	} else if score >= 10 {
		fmt.Println("Ты красавчик!")
	} else {
		fmt.Println("Надо еще многому научиться!")
	}

	isSubscribed := false
	isLogged := true

	/*
	
		and - &&
		or - ||
		not - !

	*/
	if !isSubscribed && isLogged {
		fmt.Println("Подпишись")
	} else if isSubscribed && isLogged {
		fmt.Println("Спасибо")
	} else {
		fmt.Println("Залогинся")
	}
	
}