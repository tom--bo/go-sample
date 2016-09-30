package main

import (
	"fmt"
	"github.com/lfittl/pg_query_go"
)

func main() {
	// SELECT * FROM somedb WHERE id = 100
	// SELECT * FROM otherdb WHERE uid = 1000 AND mid IS 'abcd' LIMIT 10
	ret, err := pg_query.Normalize("SELECT * FROM otherdb WHERE uid = 1000 AND name LIKE 'abcd' LIMIT 10")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", ret)

}
