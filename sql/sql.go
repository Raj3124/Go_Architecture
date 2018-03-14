package database

import(
	"fmt"
	 Conf "Go_Architecture/configuration"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//MSG "Go_Architecture/message"
)

var db *sql.DB
var err error

func Start_Con(){
	dsn:=Conf.DB_USER+":"+Conf.DB_PASS+"@"+Conf.DB_HOST+"/"+Conf.DB_NAME+"?charset=utf8"
	db,err=sql.Open("mysql",dsn)
	if err!=nil {
		fmt.Println("Cannot Connect to Database..")
	}else{
		fmt.Println("--*Starting Server *--\n")
		fmt.Println("--*Connecting to Database*--\n")
	}
}

func Stop_Con(){
	db.Close()
	//fmt.Print("Database Connection Stopped\n")
}

func Faculty(){

}

func Create(){
	_,err=db.Query("Create table newtable(id int)")
	if err!=nil{
		panic(err)
	}
}

func Insert(test string){
	fmt.Println(test)
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
