package portals

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/velebak/colly-sqlite3-storage/colly/sqlite3"
)

func IndeksOnline() {

	c := colly.NewCollector(
		colly.AllowedDomains("indeksonline.net"),
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
		7,       // Number of consumer threads
		storage, // Use sqlite queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML("div.container a[href]", func(e *colly.HTMLElement) {
		foundURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Contains(foundURL, "category") {
			return
		}
		fmt.Println("Visiting", foundURL)
		detailCollector.Visit(foundURL)
	})

	detailCollector.OnHTML("div.container", func(e *colly.HTMLElement) {
		fmt.Println("Scraping Content ", e.Request.URL.String())
		article := Article{}
		article.PortalID = 3
		article.URL = e.Request.URL.String()
		article.ArticleTitle = e.ChildText("h1.title")

		content, _ := e.DOM.Find("div.full-text").Html()
		m1 := regexp.MustCompile(`(?s)(<ins.*?</script>)`)

		article.ArticleContent = m1.ReplaceAllString(content, "")

		indekscat := e.ChildText("h3.tab_title")
		cat := strings.Fields(indekscat)

		category := cat[len(cat)-1]
		article.ArticleImage = e.ChildAttr("div.full-img > img", "src")

		if article.Category = GetCategory(categories, category); article.Category == 0 {
			return
		}

		if err := article.Save(); err != nil {
			fmt.Println("DB save error: ", err)
		}
		if err := article.SaveAPI("indeksonline"); err != nil {
			fmt.Println("Api save error: ", err)
			return
		}
	})

	q.AddURL("https://indeksonline.net/lajme/")
	q.AddURL("https://indeksonline.net/sport/")
	q.AddURL("https://indeksonline.net/showbiz/")
	q.AddURL("https://indeksonline.net/tech/")
	q.AddURL("https://indeksonline.net/kuriozitete/")
	q.AddURL("https://indeksonline.net/shendetesi/")
	q.AddURL("https://indeksonline.net/ekonomi/")
	q.Run(c)
}
