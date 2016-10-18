package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	// ファイルからの読み出し
	b, err := ioutil.ReadFile("miso_soup.jpg")
	// b []byte
	if err != nil {
		panic(err)

	}

	// ファイルへ書き込み
	// err = ioutil.WriteFile("test.jpg", b, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	webdavURL := "http://54.249.13.5:80/webdav/test.jpg"
	/* webdavへpost */
	// postWebdav(webdavURL, b)

	/* webdavから取得 */
	data := getWebdav(webdavURL)

	// mysqlへ保存

	// mysqlから取得

	// redisへ保存

	// redisから取得

}

func postWebdav(url string, data []byte) {
	buf := bytes.NewReader(data)
	// url needs protocol like http://
	req, err := http.NewRequest("PUT", url, buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	// req.Header.Add("Content-Type", content_type)
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	defer resp.Body.Close()
}

func getWebdav(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	defer resp.Body.Close()

	data, err3 := ioutil.ReadAll(resp.Body)

	if err3 != nil {
		fmt.Println(err3)
		os.Exit(1)
	}
	return data
}
