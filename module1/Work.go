package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int, 10)
	defer close(message)
	go func() {
		count := 1
		for {
			time.Sleep(1 * time.Second)
			message <- count
			count++
		}
	}()
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println(<-message)
	}
}
