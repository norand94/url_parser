package main

import (
	"fmt"
	"os"
	"url_parser/core"
)

func main(){

	if (len(os.Args) > 1) {
		fmt.Printf("url = %v \r\n"  ,os.Args[1])
		core.Parse(os.Args[1])
	} else {
		fmt.Println("Please, enter the target url")
	}
}
