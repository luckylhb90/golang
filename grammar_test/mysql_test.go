package grammar_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

type Data struct {
	rid    int
	reward string
}

var slice []Data

func Test_Insert(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	insert, err := db.Prepare("INSERT into rewards (reward , rid)values(?,?) ")
	if err != nil {
		log.Fatal(err)
	}
	res, _ := insert.Exec("Go", 1)
	fmt.Println(res)
	id, _ := res.LastInsertId()
	fmt.Println(id)
}

//golang操作mysql实现curd基础操作
func Test_Select(t *testing.T) {
	db, error := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if error != nil {
		log.Fatal(error)
	}
	//select
	rows, error := db.Query("SELECT * FROM rewards")
	for rows.Next() {
		var rid int
		var reward string
		error = rows.Scan(&rid, &reward)
		a := Data{rid, reward}
		slice = append(slice, a)
	}

	fmt.Println(slice)
	for _, v := range slice {
		fmt.Println(v)
	}
}

func Test_Update(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	update, err := db.Prepare("UPDATE rewards set reward=? where rid=?")
	res, _ := update.Exec("Java", 1)
	id, err := res.RowsAffected()
	fmt.Println(id)
}

func Test_Delete(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	del, _ := db.Prepare("DELETE from rewards where rid=?")
	res, _ := del.Exec(12)
	fmt.Println(res)
}
