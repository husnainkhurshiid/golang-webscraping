package main

import (
	"fmt"
	"scrapeQuotes/excel"
	"scrapeQuotes/model"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	var quotes []model.Quote

	// Create a callback for extracting quotes from the page
	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		quote := model.Quote{
			Text:   e.ChildText("span.text"),
			Author: e.ChildText("span small"),
			Tags:   e.ChildTexts("div.tags a.tag"),
		}
		quotes = append(quotes, quote)
	})

	// Create a callback to follow the next page link
	c.OnHTML("li.next a", func(e *colly.HTMLElement) {
		nextPageLink := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(nextPageLink))
	})

	// Start scraping on the initial page
	err := c.Visit("http://quotes.toscrape.com")

	if err != nil {
		fmt.Println("Error scraping website:", err)
		return
	}

	// export the scraped quotes
	if err := excel.ExportStructToCSV(quotes); err != nil {
		fmt.Println("Error exporting to Excel:", err)
		return
	}
}
