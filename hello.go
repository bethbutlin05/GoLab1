package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Printf("Hello from goroutine %d\n", n)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
