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

func View(sem string,sub string,w http.ResponseWriter){
var u1,u2 Student 
 rows,err=db.Query("select eid,ename from student where semester=? and eid IN (select eid from enrolled where sid IN(select sid from subject_table where sname=?))",sem,sub)
    if err!=nil{
   	 resp.DatabaseSelectionFailed(w)
    }
else{
defer rows.Close()
    if rows.Next(){
   		 rows.Scan(&u1.eid,&u1.ename)
   		 json.NewEncoder(w).Encode(u1)
}else{
             json.NewEncoder(w).Encode(u1)
            resp.Nodatafound(w)
  }
}

