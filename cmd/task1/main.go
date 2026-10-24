package main

import (
    "fmt"
    "time"
)

// printNumbers печатает числа от 1 до n с небольшими паузами.
func printNumbers(id int, n int) {
    // TODO: реализуйте цикл
	for i := 1; i <= n; i++ {
		time.Sleep(100 * time.Millisecond)
    	fmt.Printf("Worker %d: number %d\n", id, i)
	}

}

func main() {
    fmt.Println("Start lab1")
    // TODO: здесь будут вызовы printNumbers
	printNumbers(1, 5)
    fmt.Println("End lab1")
}
