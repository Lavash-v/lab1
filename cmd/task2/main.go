package main

import (
    "fmt"
    "time"
)

func printNumbers(id int, n int) {
	for i := 1; i <= n; i++ {
		time.Sleep(100 * time.Millisecond)
    	fmt.Printf("Worker %d: number %d\n", id, i)
	}

}

func main() {
    fmt.Println("Start lab1")

    go printNumbers(1, 5)
    go printNumbers(2, 5)
    go printNumbers(3, 5)
    go printNumbers(4, 5)
    go printNumbers(5, 5)

    time.Sleep(1 * time.Second) // здесь используется time.Sleep чтобы все горутины завершили свои действия
	// Если этого не сделать - программа моментально завершится, не дождавшись горутин.
	

    fmt.Println("End lab1")

}
