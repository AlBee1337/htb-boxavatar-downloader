package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/cavaliergopher/grab/v3"
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
			avatar, _ := grab.Get(os.Args[2], src)
			isTrueOneTime = false
			fmt.Printf("downloaded as %s at %s\n", avatar.Filename, os.Args[2])
		}
	})
	url := fmt.Sprintf("https://www.hackthebox.com/machines/%s", os.Args[1])
	c.Visit(url)

}
