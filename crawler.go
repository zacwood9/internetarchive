package internetarchive

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type CrawlResult struct {
	Items  []map[string]interface{}
	Count  float64
	Cursor string
	Total  float64
}

func addCursor(str string, cursor string) string {
	if cursor == "" {
		return str
	}

	index := strings.Index(str, "cursor")
	if index == -1 {
		return str + fmt.Sprintf("&cursor=%s", cursor)
	} else {
		return str[:index] + fmt.Sprintf("&cursor=%s", cursor)
	}
}

func Crawl(fields []string, collection string, stream chan CrawlResult) {
	list := strings.Join(fields, ",")
	crawlUrl := fmt.Sprintf("http://archive.org/services/search/v1/scrape?fields=%s&q=collection%%3A%s", list, collection)
	result := CrawlResult{}

	for {
		r, err := http.Get(crawlUrl)
		check(err)

		result = CrawlResult{}
		err = json.NewDecoder(r.Body).Decode(&result)
		check(err)

		if result.Cursor == "" {
			close(stream)
			break
		}

		stream <- result

		crawlUrl = addCursor(crawlUrl, result.Cursor)
	}

	//for _, item := range result.Items{
	//	date := item["date"].(string)
	//	index := strings.Index(date, "T")
	//	fmt.Println(date[:index])
	//}
	//

}
