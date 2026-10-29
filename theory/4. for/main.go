package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*	for
	
		for var := val; var condition; var action { code }
	
	*/
	for i := 0; i < 5; i++ {
		fmt.Println("Переменная i:", i)
	}

	fmt.Println("")
	// modern variation
	for i := range 5 {
		fmt.Println("Переменная i:", i)
	}

	// while in go is for without condition
	for {
		if rand.Intn(8) == 1 {
			fmt.Println("Выход из цикла")
			break
		}
		if rand.Intn(8) == 2 {
			fmt.Println("Пропуск")
			continue
		}
		fmt.Println("Продолжаем")
	}
}