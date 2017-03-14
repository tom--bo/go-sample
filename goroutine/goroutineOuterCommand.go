package main

import (
    "fmt"
    "os/exec"
    "time"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(1)
	// fmt.Println(runtime.NumCPU())
    start := time.Now()
    wg := &sync.WaitGroup{}
    c := 300

    wg.Add(c)
    for i:=0; i<c; i++ {
        go outerCommand(wg)
    }
	wg.Wait()
    end := time.Now()
    fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func outerCommand(wg *sync.WaitGroup) {
    exec.Command("sleep", "1").Run()
    wg.Done()
}
