package database

import(
	"fmt"
	"os"
	"database/sql"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
Conf "Go_Architecture/configuration"
resp "Go_Architecture/response"
	//MSG "Go_Architecture/message"
)

var db *sql.DB
var err error

func Start_Con(){
	dsn:=Conf.DB_USER+":"+Conf.DB_PASS+"@"+Conf.DB_HOST+"/"+Conf.DB_NAME+"?charset=utf8"
	db,err=sql.Open("mysql",dsn)
	if err!=nil {
		fmt.Println("Cannot Connect to Database..")
		os.Exit(3)
	}else{
		fmt.Println("--*Started Server *--\n")
		fmt.Println("--*Connected to Database*--\n")
	}
	err = db.Ping()
if err != nil {
	fmt.Println("Cannot Connect to Database")
	os.Exit(3)
}
}

func Stop_Con(){
	db.Close()
	//fmt.Print("Database Connection Stopped\n")
}

func Insert(test string,w http.ResponseWriter){
	_,err=db.Query("insert into login(test) values(?)",test)
	if err!=nil{
		resp.DatabaseInsertionFailed(w)
	}
}
