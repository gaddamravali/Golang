package main

import "fmt"

func EvenOdd() {

}

func main() {
	var s = []int{1, 2, 3, 4, 5, 6, 22, 23}
	chOdd := make(chan int)
	chEven := make(chan int)

	go Even(chEven)
	go Odd(chOdd)

	for i := 0; i < len(s); i++ {
		if s[i]%2 != 0 {
			chOdd <- s[i]
		} else {
			chEven <- s[i]
		}
	}
}

func Odd(ch <-chan int) {
	for s := range ch {
		fmt.Println("Odd number :", s)
	}
}

func Even(ch <-chan int) {
	for s := range ch {
		fmt.Println("Even number :", s)
	}
}
