package main

import (
    "fmt"
    "time"
)

func printNumbers(n int) {
	for i := 0; i <= n; i++ {
		go func() {
		fmt.Printf("Gotoutine %d says HIIIIIIIIIII\n", i)
		}()
	}

}

func main() {
    fmt.Println("Start lab1")

    printNumbers(1000)

	time.Sleep(500 * time.Millisecond)

    fmt.Println("End lab1")

}
