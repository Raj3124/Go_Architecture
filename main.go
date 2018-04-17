package main

import(
	DB "Go_Architecture/sql"
	"fmt"
	API "Go_Architecture/api"
	"net/http"
)

func main(){
	DB.Start_Con()
	http.HandleFunc("/",API.ErrorMessage)
	http.HandleFunc("/insert",API.Insert)
	err:=http.ListenAndServe(":2005",nil)
	if err!=nil{
		fmt.Println("Cannot Start Server... check IP and port number")
	}
	//DB.Stop_Con()
}


