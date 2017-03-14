package main

import (
    "fmt"
    "time"
    "sync"
)

var test []string

func main() {
    ch := make(chan string)
    wg := new(sync.WaitGroup)

    go say("1: ", ch, wg)
    go say("2: ", ch, wg)
    go say("3: ", ch, wg)
    time.Sleep(1 * time.Second)
    wg.Add(5)
    
    tslice := []string{"first", "second"}
    test = append(test, tslice...)
    ch <- "test 1"
    ch <- "test 2"
    ch <- "test 3"
    ch <- "test 4"
    ch <- "test 5"
    wg.Wait()
    fmt.Println("-----")
    for _, m := range test {
        fmt.Println(m)
    }
}

func say(num string, ch chan string, wg *sync.WaitGroup) {
    for ;; {
        message := <-ch
        fmt.Println(num + message)
        // mutex使わないと時々欠損する
        test = append(test, message)
        wg.Done()
        time.Sleep(10 * time.Millisecond)
    }
}

