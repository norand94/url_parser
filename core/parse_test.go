package core

import (
	"testing"
	"github.com/PuerkitoBio/goquery"
	"container/list"
	"strings"
)

func TestParse(t *testing.T) {
	doc, err := goquery.NewDocument("https://www.google.com")
	if err != nil {
		t.Error(err.Error())
	}
	l := Parse(doc)
	checkList := list.New()
	checkList.PushBack("https://www.google.ru/imghp?hl=ru&tab=wi")
	checkList.PushBack("https://maps.google.ru/maps?hl=ru&tab=wl")
	checkList.PushBack("https://play.google.com/?hl=ru&tab=w8")

	if !check(checkList, l) {
		t.Error("Отсутствует необходимые урл!")
	}

	checkList = list.New()
	checkList.PushBack("https::/yandex.ru")
	if check(checkList, l) {
		t.Error("Данного урл там быть не должно")
	}

}

func check(short, long *list.List) bool {
	l_el := long.Front()
	s_el := short.Front()
	for  s_el != nil && l_el != nil {
		if !strings.EqualFold(l_el.Value.(string), s_el.Value.(string)) {
			return false
		}
		l_el = l_el.Next()
		s_el = s_el.Next()
	}
	return true
}