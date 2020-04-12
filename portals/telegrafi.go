package portals

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/velebak/colly-sqlite3-storage/colly/sqlite3"
)

func Telegrafi() {
	c := colly.NewCollector(
		colly.AllowedDomains("telegrafi.com"),
	)

	storage := &sqlite3.Storage{
		Filename: "./results.db",
	}

	defer storage.Close()

	detailCollector := c.Clone()

	err := detailCollector.SetStorage(storage)

	if err != nil {
		panic(err)
	}

	q, _ := queue.New(
		6, // Number of consumer threads
		storage, // Use sqlite queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
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
		m1 := regexp.MustCompile(`(?s)(<iframe sandbox=".*?</iframe>)`)
		article.ArticleContent = m1.ReplaceAllString(content, "")

		category := e.ChildText("a.article-category")
		article.ArticleImage = e.ChildAttr("div.featured-image > figure > img", "src")

		if article.Category = GetCategory(categories, category); article.Category == 0 {
			return
		}

		if err := article.Save(); err != nil {
			fmt.Println("DB save error: ", err)
		}
		if err := article.SaveAPI("telegrafi"); err != nil {
			fmt.Println("Api save error: ", err)
			return
		}
	})

	q.AddURL("https://telegrafi.com/lajme/")
	q.AddURL("https://telegrafi.com/sport/")
	q.AddURL("https://telegrafi.com/magazina/")
	q.AddURL("https://telegrafi.com/shendetesi/")
	q.AddURL("https://telegrafi.com/teknologji/")
	q.AddURL("https://telegrafi.com/fun/")
	q.Run(c)
}
