package main

import (
	"fmt"
	"os"
)

func main(){

	if (len(os.Args) > 1) {
		fmt.Printf("url = %v"  ,os.Args[1])
	} else {
		fmt.Println("Please, enter the target url")
	}
}
