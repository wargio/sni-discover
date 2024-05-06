package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

type Discover struct{}

func collyErrorHandler(r *colly.Response, err error) {
	log.Fatal("Error:", err)
}

func (disco *Discover) Run(config *Config, rsv *Resolver) {
	wg := &sync.WaitGroup{}
	c := colly.NewCollector()
	c.OnError(collyErrorHandler)
	c.OnHTML("td.nowrap > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		url := fmt.Sprintf("https://bgp.he.net%s#_dnsrecords", link)
		wg.Add(1)
		disco.handleLink(url, rsv, wg)
	})

	url := fmt.Sprintf("https://bgp.he.net/ip/%s", config.target)
	err := c.Visit(url)
	if err != nil {
		log.Fatal("Failed to get DNS records", err)
	}
	wg.Wait()
}

func (disco *Discover) handleLink(url string, rsv *Resolver, wg *sync.WaitGroup) {
	snis := map[string]bool{}

	c := colly.NewCollector()
	c.OnError(collyErrorHandler)
	c.OnHTML("tr", func(row *colly.HTMLElement) {
		i := 0
		row.ForEach("td", func(_ int, e *colly.HTMLElement) {
			text := trim(e.Text)
			if i == 2 && len(text) > 0 {
				tokens := strings.Split(text, ", ")
				for _, sni := range tokens {
					sni = removeNs(sni)
					if !isSNI(sni) {
						continue
					}
					if _, ok := snis[sni]; ok {
						continue
					}
					snis[sni] = true
					rsv.AddSNI(sni)
				}
			}
			i++
		})
	})
	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal("Failed to get DNS records", err)
	}
}
