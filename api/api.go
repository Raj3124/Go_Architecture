package api

import(
	"net"
	"fmt"
	"net/http"
	"encoding/json"
	DB "Go_Architecture/sql"
	resp "Go_Architecture/response"
	logs "Go_Architecture/logs"
	reqs "Go_Architecture/request"
)

func ErrorMessage(w http.ResponseWriter,r *http.Request){
	resp.Sendresponse(w, r.Body)
}

func Insert(w http.ResponseWriter,r *http.Request){
	if(r.Method == "POST"){
		// fmt.Println()
		decoder := json.NewDecoder(r.Body)
		var t reqs.InsertStruct
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println("error in decoding")
		}
		defer r.Body.Close()
		if t.Test == "" {
		fmt.Println("empty")
		}else{
			DB.Insert(t.Test, w)
		}
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
