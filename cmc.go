package main

import (
	"log"
	"github.com/gocolly/colly"
	"strings"
)

func main() {
	c := colly.NewCollector()
	header := "Name, Market Cap, Price, Change (24h)"
	result := make([]string, 0)
	result = append(result, header)
	c.OnHTML("#currencies tbody tr", func(e *colly.HTMLElement) {
		values := []string{
			e.ChildText(".currency-name-container"),
			e.ChildText(".market-cap"),
			e.ChildText(".price"),
			e.ChildText(".percent-change"),
		}
		slice := strings.Join(values, ", ")
		log.Println(slice)
		result = append(result, slice)
	})
	//log.Println(strings.Join(result, "\n"))
	c.Visit("https://coinmarketcap.com")
	log.Println("All done!")
}