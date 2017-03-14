package main

import (
	"fmt"
	"sync"
	// "runtime"
	// "time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go calc(wg)
	}
	fmt.Println("before wait")
	wg.Wait()
	fmt.Println("ok")
}

func calc(wg *sync.WaitGroup) {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
		if i > 40 {
			i -= 20
		}
		sum /= 3000
	}
	wg.Done()
}
