package main

import (
	"fmt"
	"os"
	"url_parser/core"
	"github.com/PuerkitoBio/goquery"
)

func main(){

	if (len(os.Args) > 1) {
		doc, err := goquery.NewDocument(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		list := core.Parse(doc)
		core.Print(list)
	} else {
		fmt.Println("Please, enter the target url")
	}
}
