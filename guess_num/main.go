package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	foo()
}

func foo() {
	for {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(100) + 1

		var input string
		fmt.Println("Hello, please input a guess number")
	loop:
		fmt.Scanln(&input)

		if input == "stop" {
			return
		}

		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("input error: ", err)
		}

		switch {
		case i < randNum:
			fmt.Println("less than the number, try again")
			goto loop
		case i > randNum:
			fmt.Println("greater than the number, try again")
			goto loop
		case i == randNum:
			fmt.Println("great! you did the right choice")
		}
	}
}
