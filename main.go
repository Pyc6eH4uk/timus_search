package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func search(nickname string) {
	nickname = strings.Replace(nickname, " ", "+", -1)
	var row []string
	var result [][]string
	response, err := http.Get("http://acm.timus.ru/search.aspx?Str=" + nickname)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}
	var responseNickname string
	doc.Find("table").Each(func(_ int, table *goquery.Selection) {
		contents, _ := table.Attr("class")
		if contents == "ranklist" {
			table.Find("tr").Each(func(_ int, tr *goquery.Selection) {
				tr.Find("td").Each(func(_ int, td *goquery.Selection) {
					td.Find("a").Each(func(_ int, a *goquery.Selection) {
						responseNickname = a.Text()
						if responseNickname == strings.Replace(nickname, "+", " ", -1) {
							row = append(row, td.Text())
							result = append(result, row)
						}
					})
					row = nil
					if !strings.Contains(responseNickname, td.Text()) {
						if responseNickname == strings.Replace(nickname, "+", " ", -1) {
							row = append(row, td.Text())
							result = append(result, row)
						}
					}
				})
			})
		}
	})
	fmt.Println("result = ", result)
}

func main() {
	// insert nickname from timus
	search("")
}
