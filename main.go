package main

import "fmt"

func main() {
	oddch := make(chan int)
	evench := make(chan int)

	go PrintNum(oddch, 1)
	go PrintNum(evench, 2)

	for i := 0; i < 100; i++ {
		select {
		case num := <-oddch:
			fmt.Println(num)
		case num := <-evench:
			fmt.Println(num)

		}
	}
}

func PrintNum(ch chan int, start int) {
	for i := start; i <= 100; i++ {
		ch <- i
	}
}
