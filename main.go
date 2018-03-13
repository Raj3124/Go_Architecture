package main

import(
	DB "../myarch/sql"
	MSG "../myarch/message"
	API "../myarch/api"
	"net/http"
)

func main(){
	DB.Start_Con()
	http.HandleFunc("/",API.ErrorMessage)
	http.HandleFunc("/insert",API.Insert)
	http.HandleFunc("/delete",API.Delete)
	http.HandleFunc("/update",API.Update)
	http.HandleFunc("/create",API.Create)
	err:=http.ListenAndServe(":2008",nil)
	if err!=nil{
		MSG.DB_message(err)
	}
	DB.Stop_Con()
}
