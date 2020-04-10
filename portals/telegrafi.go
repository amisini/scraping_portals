package portals

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/microcosm-cc/bluemonday"
)

func Telegrafi() {
	c := colly.NewCollector(
		colly.AllowedDomains("telegrafi.com"),
		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./telegrafi_cache"),
	)

	detailCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML("div.aktuale-widget a[href]", func(e *colly.HTMLElement) {
		foundURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Contains(foundURL, "category") {
			return
		}
		fmt.Println("Visiting", foundURL)
		detailCollector.Visit(foundURL)
	})

	detailCollector.OnHTML("div.article-container", func(e *colly.HTMLElement) {
		fmt.Println("Scraping Content ", e.Request.URL.String())
		article := Article{}
		article.PortalID = 1
		article.URL = e.Request.URL.String()
		article.ArticleTitle = e.ChildText("h1")

		content, _ := e.DOM.Find("div.article-body").Html()
		p := bluemonday.UGCPolicy()
		article.ArticleContent = p.Sanitize(content)

		category := e.ChildText("a.article-category")
		article.ArticleImage = e.ChildAttr("div.featured-image > figure > img", "src")

		if article.Category = GetCategory(categories, category); article.Category == 0 {
			return
		}

		if err := article.Save(); err != nil {
			fmt.Println("DB save error: ", err)
			return
		}
		if err := article.SaveAPI("telegrafi"); err != nil {
			fmt.Println("Api save error: ", err)
			return
		}
	})

	c.Visit("https://telegrafi.com/lajme/")
	c.Visit("https://telegrafi.com/sport/")
	c.Visit("https://telegrafi.com/magazina/")
	c.Visit("https://telegrafi.com/shendetesi/")
	c.Visit("https://telegrafi.com/teknologji/")
	c.Visit("https://telegrafi.com/fun/")
}
