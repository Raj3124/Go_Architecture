package message

import(
	"fmt"
)

func Err_message(err error){
	fmt.Print("You have got error\n")
	panic(err)
	fmt.Print("Try again Later\n")
}

func DB_message(err error){
	fmt.Print("Database not able to connect due to ")
	panic(err)
	fmt.Print("Try again ")
}
