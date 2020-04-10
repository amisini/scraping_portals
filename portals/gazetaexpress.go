package portals

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/microcosm-cc/bluemonday"
)

func GazetaExpress() {

	c := colly.NewCollector(
		colly.AllowedDomains("gazetaexpress.com", "www.gazetaexpress.com"),
		colly.UserAgent(userAgent),
		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./gazetaexpress_cache"),
	)

	detailCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML("div.row a[href]", func(e *colly.HTMLElement) {
		foundURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Contains(foundURL, "category") {
			return
		}
		fmt.Println("Visiting", foundURL)
		detailCollector.Visit(foundURL)
	})

	detailCollector.OnHTML("div.single__content", func(e *colly.HTMLElement) {
		fmt.Println("Scraping Content ", e.Request.URL.String())
		article := Article{}
		article.PortalID = 1
		article.URL = e.Request.URL.String()
		article.ArticleTitle = e.ChildText("h2.single__title")

		content, _ := e.DOM.Html()
		m1 := regexp.MustCompile(`(?s)(<h2.*?</div>)`)
		myContent := m1.ReplaceAllString(content, "")

		p := bluemonday.UGCPolicy()
		article.ArticleContent = p.Sanitize(myContent)

		category := e.ChildText("div.single__author h2")
		article.ArticleImage = e.ChildAttr("figure > img", "src")

		if article.Category = GetCategory(categories, category); article.Category == 0 {
			return
		}

		if err := article.Save(); err != nil {
			fmt.Println("DB save error: ", err)
			return
		}
		if err := article.SaveAPI("gazetaexpress"); err != nil {
			fmt.Println("Api save error: ", err)
			return
		}
	})

	c.Visit("https://www.gazetaexpress.com/lajme/")
	c.Visit("https://www.gazetaexpress.com/sport/")
	c.Visit("https://www.gazetaexpress.com/roze/")
	c.Visit("https://www.gazetaexpress.com/shneta/")
}
