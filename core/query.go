package core

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func Parse (url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(doc)
}
