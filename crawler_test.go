package internetarchive

import (
	"testing"
)

func TestCrawl(t *testing.T) {
	stream := make(chan CrawlResult)
	go Crawl([]string{"date", "num_reviews"}, "gratefuldead", stream)

	for range stream {

	}
}
