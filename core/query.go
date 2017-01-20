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
	fmt.Println("links:")
	doc.Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		text := a.Text()
		href, ok := a.Attr("href")
		if ok {
			fmt.Printf("%s  -  %s", href, text)
		}
	})

}
