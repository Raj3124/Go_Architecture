package api

import(
	"net"
	"fmt"
	"net/http"
	DB "myarch/sql"
	resp "myarch/response"
	logs "myarch/logs"
)

func ErrorMessage(w http.ResponseWriter,r *http.Request){

	resp.Sendresponse(w, r.Body)
}

func Insert(w http.ResponseWriter,r *http.Request){
	if(r.Method == "POST"){

	fmt.Fprintln(w,"Inserting Data"+r.Method)
	DB.Insert()
}else{
	ip, port, err2 := net.SplitHostPort(r.RemoteAddr)
	if err2 != nil{
		w.Write([]byte("Cannot get IP"))
	}
	userIP := net.ParseIP(ip)
	logs.Create_log(port,userIP)
	resp.InvalidMethod(w, r.Body)
}

}

func Delete(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Deleting data"+r.Method)
	DB.Delete()
}

func Update(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Updating data"+r.Method)
	DB.Update()
}

func Create(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Creating new Function"+r.Method)
	DB.Create()
}
