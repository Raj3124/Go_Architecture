package database

import(
	"fmt"
	 Conf "myarch/configuration"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	MSG "myarch/message"
)

var db *sql.DB
var err error

func Start_Con(){
	dsn:=Conf.DB_USER+":"+Conf.DB_PASS+"@"+Conf.DB_HOST+"/"+Conf.DB_NAME+"?charset=utf8"
	db,err=sql.Open("mysql",dsn)
	if err!=nil {
		MSG.Err_message(err)
	}else{
		fmt.Print("Connection Started\n")
	}
}

func Stop_Con(){
	db.Close()
	fmt.Print("Connection Stopped\n")
}

func Create(){
	_,err=db.Query("Create table newtable(id int)")
	if err!=nil{
		panic(err)
	}
}

func Insert(){
	//_,err=db.Query("Insert into newtable values(2)")
	//if err!=nil{
//		panic(err)
//	}
}

func Update(){
	//_,err=db.Query("Update newtable set id=1 where id =2")
	//if err!=nil{
//		panic(err)
//	}
}

func Delete(){
	//_,err=db.Query("delete from newtable where id=1")
	//if err!=nil{
//		panic(err)
//	}
}
