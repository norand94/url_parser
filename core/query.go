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
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		//text := a.Text()
		href, ok := s.Attr("href")
		if ok {
			fmt.Printf("%s \n", href)
		}
	})

}
