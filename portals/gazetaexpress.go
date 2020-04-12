package portals

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/velebak/colly-sqlite3-storage/colly/sqlite3"
)

func GazetaExpress() {

	c := colly.NewCollector(
		colly.AllowedDomains("gazetaexpress.com", "www.gazetaexpress.com"),
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
		4,       // Number of consumer threads
		storage, // Use sqlite queue storage
	)

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
		article.PortalID = 2
		article.URL = e.Request.URL.String()
		article.ArticleTitle = e.ChildText("h2.single__title")

		content, _ := e.DOM.Html()
		m1 := regexp.MustCompile(`(?s)(<h2.*?</div>)`)
		removeHeader := m1.ReplaceAllString(content, "")
		m2 := regexp.MustCompile(`(?s)(<ins.*?</script>)`)
		myContent := m2.ReplaceAllString(removeHeader, "")

		article.ArticleContent = myContent

		category := e.ChildText("div.single__author h2")
		article.ArticleImage = e.ChildAttr("figure > img", "src")

		if article.Category = GetCategory(categories, category); article.Category == 0 {
			return
		}

		if err := article.Save(); err != nil {
			fmt.Println("DB save error: ", err)
		}
		if err := article.SaveAPI("gazetaexpress"); err != nil {
			fmt.Println("Api save error: ", err)
			return
		}
	})

	q.AddURL("https://www.gazetaexpress.com/lajme/")
	q.AddURL("https://www.gazetaexpress.com/sport/")
	q.AddURL("https://www.gazetaexpress.com/roze/")
	q.AddURL("https://www.gazetaexpress.com/shneta/")
	q.Run(c)
}
