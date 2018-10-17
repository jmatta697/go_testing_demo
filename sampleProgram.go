package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	fmt.Println("** Sample program demonstrating testing in Go **")
	//userInputInt := GetIntFromUser()
	//fmt.Println(Calculate(userInputInt))
	//fmt.Println(Fib(userInputInt))

	counter := 0
	for {
		if counter == 10 {
			break
		}
		fmt.Println(PickRandNumber())
		counter++
	}

}


func Calculate(x int) int {
	result := x + 2
	return result
}

// calculates the nth fibonacci number
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func PickRandNumber() int {
	return rand.Intn(10) + 1
}


// -----------------------------------------------------

func GetIntFromUser() int {
	fmt.Print("Enter an integer: ")
	var userInt int
	_, err := fmt.Scanf("%d", &userInt)
	if err != nil {
		log.Print(err)
	}
	return userInt
}

