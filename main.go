package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, value := range s {
		sum += value
	}

	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	channel := make(chan int)

	go sum(s[:len(s)/2], channel)
	go sum(s[len(s)/2:], channel)

	x, y := <-channel, <-channel

	fmt.Println(x, y, x+y)
}
