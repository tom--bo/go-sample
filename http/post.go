package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go geturl("http://localhost:8080/sleep/2", wg)
	wg.Add(1)
	go geturl("http://localhost:8080/sleep/1", wg)
	wg.Wait()
}

func geturl(url string, wg *sync.WaitGroup) string {
	defer wg.Done()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}
