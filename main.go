package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type httpProxy struct {
	ip   string
	port int
}

func scrapeProxyDaily() (proxies []httpProxy, err error) {
	var rawProxies []string

	c := colly.NewCollector(
		colly.AllowedDomains("proxy-daily.com"),
		colly.MaxDepth(1),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("div.centeredProxyList.freeProxyStyle", func(e *colly.HTMLElement) {
		rawProxies = append(rawProxies, e.Text)
	})

	c.Visit("https://proxy-daily.com")

	return
}

func scrapeProxies(site string) (proxies []httpProxy, err error) {
	switch site {
	case "proxyDaily":
		return scrapeProxyDaily()
	default:
		return
	}
}

func main() {
	proxies, err := scrapeProxies("proxyDaily")
	if err != nil {
		panic(err)
	}
	for _, p := range proxies {
		fmt.Println(p)
	}
}
