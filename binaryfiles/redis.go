package main

import (
	// "fmt"
	redis "gopkg.in/redis.v5"
	"io/ioutil"
)

func main() {
	remoteIP := "54.199.151.194"

	// ファイルからの読み出し
	// b []byte
	b := readFromFile("miso_soup.jpg")

	// ファイルへ書き込み
	// writeFile("test.jpg", b)

	///////////////////
	/* redisへの接続 */
	///////////////////
	var rd *redis.Client
	rd = redis.NewClient(&redis.Options{
		Addr: remoteIP + ":6379",
		DB:   0,
	})

	// redisへ保存 （[]byteを文字列にして突っ込む)
	keyName := "sample1"
	redis_b := string(b)
	rd.Set(keyName, redis_b, 0)

	// redisから取得
	redisData, _ := rd.Get(keyName).Result()
	// 必要であれば[]byteに書き換える
	writeFile("redis_result.jpg", []byte(redisData))
}

func readFromFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	// b []byte
	if err != nil {
		panic(err)
	}
	return b
}

func writeFile(filename string, data []byte) {
	err := ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		panic(err)
	}
}
