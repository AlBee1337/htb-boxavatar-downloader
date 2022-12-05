package main

import (
	"fmt"
	"regexp"

	"github.com/cavaliercoder/grab/v3"
	"github.com/gocolly/colly/v2"
)

func main() {
	src := ""
	c := colly.NewCollector()
	isTrue := false
	isTrueOneTime := true
	c.OnHTML("img", func(e *colly.HTMLElement) {
		src = e.Attr("src")
		isTrue, _ = regexp.MatchString("https://www.hackthebox.com/storage/avatars/[[:ascii:]]+.png", src)
		if isTrue && isTrueOneTime {
			// fmt.Printf(src)
			donwloadedAvatar, _ := grab.Get(".", src)
			isTrueOneTime = false
			fmt.Println(donwloadedAvatar)
		}
		// fmt.Println(src)
	})
	c.Visit("https://www.hackthebox.com/machines/steamcloud")
	// fmt.Println(src)

	// match, _ := regexp.MatchString("peach", "p([a-z]+)ch")
	// fmt.Println(match)
}
