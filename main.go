package main

import (
	"fmt"
	"time"
)

func main() {
	nums := make(chan int)
	squares := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		nums <- 3
	}()

	go func() {
		time.Sleep(2 * time.Second)
		squares <- 3 * 3
	}()

	for i := 0; i < 2; i++ {
		select {
		case num := <-nums:
			fmt.Println(num)
		case square := <-squares:
			fmt.Println(square)
		}
	}
}
