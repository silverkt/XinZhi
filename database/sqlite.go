package main;
import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	)
func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	//创建数据库表
createsql :=	`CREATE TABLE userinfo
	(
	username varchar(255),
	departname varchar(255),
	created varchar(255)
	);`


	stmt, err := db.Prepare(createsql)
checkErr(err)

_ , err = stmt.Exec()
checkErr(err)

//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
checkErr(err)

_ , err = stmt.Exec("astaxie", "研发部门", "2012-12-09")
checkErr(err)



//读取数据
rows, err := db.Query("select * from userinfo")
checkErr(err)




for rows.Next() {

    var username string
    var department string
    var created string
    err = rows.Scan( &username, &department, &created)
    checkErr(err)
 
    fmt.Println(username)
    fmt.Println(department)
    fmt.Println(created)
}

db.Close();


}

func checkErr(err error) {
	if err != nil {
	panic(err)
	}
	}