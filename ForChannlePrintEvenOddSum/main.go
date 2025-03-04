package main

import "fmt"

func EvenSum(evenChan chan int, slice []int) {
	sum := 0

	for i := 0; i < len(slice); i++ {
		if i%2 == 0 {
			sum += slice[i]
		}
	}

	evenChan <- sum
}

func OddSum(oddChan chan int, slice []int) {
	sum := 0

	for i := 0; i < len(slice); i++ {
		if i%2 != 0 {
			sum += slice[i]
		}
	}

	oddChan <- sum
}

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenChan := make(chan int)
	oddChan := make(chan int)

	go EvenSum(evenChan, slice)
	go OddSum(oddChan, slice)

	even := <-evenChan
	odd := <-oddChan

	fmt.Println("Even: ", even)
	fmt.Println("Odd: ", odd)

}
