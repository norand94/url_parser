package core

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"container/list"
)


func Parse (doc *goquery.Document) *list.List {
	links := list.New()
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		//text := a.Text()
		href, ok := s.Attr("href")
		if ok {
			links.PushBack(href)
		}
	})
	return links
}

func Print(l *list.List)  {
	fmt.Println("links:")
	for e := l.Front(); e != nil; e = e.Next()  {
		//str := strings.Replace(e.Value.(string), "\n", "\\n", -1)
		fmt.Println(e.Value)
	}
}