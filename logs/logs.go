package log

import(
	"fmt"
	"os"
	"net"
	"time"

)

func Create_log(port string ,userIP net.IP){

	file, err1 := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err1 != nil {
	fmt.Println("error")
		}
	fmt.Fprint(file,"Invalid method ip:",userIP)
	fmt.Fprint(file,"port:",port)
	fmt.Fprint(file,"time:",time.Now())
	fmt.Fprintln(file,"--------")
	file.Close()
}
