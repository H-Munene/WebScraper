package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main(){
	
	c:= colly.NewCollector(
		colly.AllowedDomains("strathmore.edu"),
	)
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping: ",r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something bogus happened while trying to access ",r.Request.URL,"__",err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status code is:", r.StatusCode)
	})

	c.OnHTML("title", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit("https://strathmore.edu/")
}

