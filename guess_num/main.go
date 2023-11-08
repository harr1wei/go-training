package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var interrupt = make(chan os.Signal)

func main() {
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)
	go foo()
	select {
	case <-interrupt:
		return
	}
}

func foo() {
	for {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(100) + 1

		fmt.Println("Hello, please input a guess number")
	loop:
		input := ""
		fmt.Scanln(&input)
		if input == "stop" {
			return
		}

		i, err := strconv.Atoi(input)
		if err != nil {
			goto loop
		}

		switch {
		case i < randNum:
			fmt.Println("less than the number, try again")
			goto loop
		case i > randNum:
			fmt.Println("greater than the number, try again")
			goto loop
		case i == randNum:
			fmt.Printf("great! you did the right choice, number is %d\n\n", i)
		}
	}
}
