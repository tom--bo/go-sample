package main

import (
	"fmt"
	"sync"
	// "runtime"
	// "time"
)

var cnt = 0

func main() {
	// fmt.Println(runtime.NumCPU())
	fmt.Println(cnt)
	m := new(sync.Mutex)
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go sum(1, m, wg)
	}
	wg.Wait()
	fmt.Println(cnt)
}

func sum(delta int, m *sync.Mutex, wg *sync.WaitGroup) {
	m.Lock()
	defer m.Unlock()
	c := cnt
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += delta
	}
	c += sum
	cnt = c
	wg.Done()
}
