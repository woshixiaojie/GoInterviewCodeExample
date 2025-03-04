package main

import (
	"fmt"
	"time"
)

// 1 - n 个goroutine
var worker = 5

func main() {

	// 创建n个 channel
	slice := make([]chan int, worker)
	for i := 0; i < worker; i++ {
		slice[i] = make(chan int)
	}

	// 创建n个 goroutine
	for i := 0; i < worker; i++ {

		// 开启goroutine
		go Print(slice[i], i)
	}

	number := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < worker; j++ {

			// 往n个channel里，循环发送，不断+1的数字
			slice[j] <- number
			number++

			// 休息1s
			time.Sleep(1 * time.Second)
		}
	}
}

func Print(ch chan int, workerID int) {
	for {
		fmt.Printf("goroutine %d 打印数字 %d \n ", workerID, <-ch)
	}
}
