package main

import (
	"fmt"
	"strings"

	"github.com/amisini/scraping_portals/portals"
	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("telegrafi.com"),
		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./telegrafi_cache"),
	)

	detailCollector := c.Clone()

	allArticles := []portals.Article{}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
		foundURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Contains(foundURL, "category") {
			return
		}
		fmt.Println("Visiting", foundURL)
		detailCollector.Visit(foundURL)
	})

	detailCollector.OnHTML(`div.article-container`, func(e *colly.HTMLElement) {
		fmt.Println("Scraping Content ", e.Request.URL.String())
		article := portals.Article{}
		article.URL = e.Request.URL.String()
		article.ArticleTitle = e.ChildText("h1")
		article.ArticleContent = e.ChildText("div.article-body")
		category := e.ChildText("a.article-category")
		article.ArticleImage = e.ChildAttr("div.featured-image > figure > img", "src")

		categories := map[string]int8{
			"lajme":      1,
			"sport":      2,
			"magazina":   3,
			"teknologji": 4,
			"fun":        5,
			"shendetesi": 6,
			"ekonomi":    7,
		}

		if article.Category = GetCategory(categories, category); article.Category == 0 {
			return
		}

		if err := article.Save(); err != nil {
			fmt.Println("DB save error: ", err)
		}
		if err := article.SaveAPI(); err != nil {
			fmt.Println("Api save error: ", err)
		}
		allArticles = append(allArticles, article)
	})

	c.Visit("https://telegrafi.com/")
}

func GetCategory(categories map[string]int8, cat string) int8 {
	for key, value := range categories {
		if strings.Contains(strings.ToLower(cat), key) {
			return value
		}
	}
	return 0
}
