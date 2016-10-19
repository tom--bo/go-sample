package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

func main() {
	remoteIP := "54.199.151.194"

	// ファイルからの読み出し
	// b []byte
	b := readFromFile("miso_soup.jpg")

	///////////////////
	/* mysqlへの接続 */
	///////////////////

	// sqlxを使った接続（pixiv社内ISUCONの実装を参考）
	user := "isucon"
	password := "isucon"
	host := remoteIP
	port := "3306"
	dbname := "mysample"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	// var db *sqlx.DB
	db, merr := sqlx.Open("mysql", dsn)
	if merr != nil {
		fmt.Println(merr.Error())
		return
	}
	defer db.Close()

	/* Table definition */
	// CREATE TABLE binary_table (
	// 	id int NOT NULL auto_increment,
	// 	blobdata MEDIUMBLOB,
	// 	PRIMARY KEY(id)
	// )

	// mysqlへ保存
	// blob型に入れる場合でもstringにcastする
	mydata := string(b)
	pid := insertDataToMySQL(db, mydata)
	fmt.Printf("Last inserted id is %d\n", pid)

	// mysqlから取得
	mygetdata := getDataFromMySQL(db, pid)

	writeFile("mysqldata.jpg", mygetdata)

}

func getDataFromMySQL(db *sqlx.DB, id int) []byte {
	var data []byte
	err := db.Get(&data, "SELECT blobdata FROM `binary_table` WHERE `id` = ?", id)
	if err != nil {
		return nil
	}
	return data
}

func insertDataToMySQL(db *sqlx.DB, data string) int {
	query := "INSERT INTO `binary_table` (`blobdata`) VALUES (?)"
	result, eerr := db.Exec(
		query,
		data,
	)

	if eerr != nil {
		fmt.Println(eerr.Error())
		return -1
	}

	pid, lerr := result.LastInsertId()
	if lerr != nil {
		fmt.Println(lerr.Error())
		return -1

	}
	return int(pid)
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
