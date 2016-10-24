package main

import (
    "fmt"
    "os/exec"
    "time"
    "runtime"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
    start := time.Now()
    go outerCommand()
    go outerCommand()
    go outerCommand()
    go outerCommand()
    go outerCommand()
    end := time.Now()
    fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func outerCommand() {
    exec.Command("sleep", "1").Run()
}
