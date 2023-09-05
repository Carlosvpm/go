package main

import (
	"fmt"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n ", workerId, x)
	}

}

// go routine
// Multi Thread usando go func()
// Para cada função foi criada uma thread de execução.
func main() { // T0
	n := 1000000
	channel := make(chan int)
    for i := 0; i <= n; i++ {
		go worker(i, channel) // T1
	}
	

	for i := 0; i <= n; i++ {
		channel <- i
	}
}
