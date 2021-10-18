package sql

import (
	"database/sql"
	"fmt"
	"log"
)

type dbInfo struct {
	user     string
	pwd      string
	url      string
	engine   string
	database string
}

var db1 = dbInfo{"root", "mypassword", "localhost:7545", "mysql", "test"}

//var query 중

func dbQuery(db dbInfo, query string) (count int) {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	return count
}

func main() {
	result := dbQuery(db1, query)
	print(result)
}
