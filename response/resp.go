package response

import(

  "encoding/json"
  "net/http"
  config "Go_Architecture/configuration"
  "io"
)
type ErrorResponses struct {
    Msg string `json:"message"`
    Code string `json:"code"`
}

func Sendresponse(w http.ResponseWriter,r io.ReadCloser){
  w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(500)
  responses := &ErrorResponses{
    Msg : config.INVALID_FORMAT,
    Code : config.INTERNAL_SERVER_ERROR}
  responses1, _ := json.Marshal(responses)
  w.Write([]byte(responses1))
}

func InvalidMethod(w http.ResponseWriter,r io.ReadCloser){
  w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(405)
  responses := &ErrorResponses{
    Msg : config.INVALID_METHOD_MESSAGE,
    Code : config.INVALID_METHOD_CODE}
  responses1, _ := json.Marshal(responses)
  w.Write([]byte(responses1))
}

func ServerError(){
  
}
