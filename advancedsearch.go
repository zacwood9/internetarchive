package internetarchive

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type advancedSearch struct {
	Query   Query
	Options Options
}

type Query struct {
	Collection string
	SearchTerm string
}

type Options struct {
	Fields []string
	SortBy string
	Rows   int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func constructUrl(query Query, options Options) string {
	start := "https://archive.org/advancedsearch.php?q="

	// format collection and query string
	end := fmt.Sprintf("collection:\"%s\" %s", query.Collection, query.SearchTerm)
	end += "&"

	// append all fields
	for i := 0; i < len(options.Fields); i++ {
		end += fmt.Sprintf("fl[]=%s&", options.Fields[i])
	}

	// append sort method
	end += fmt.Sprintf("sort[]=%s&", options.SortBy)
	end += "sort[]=&sort[]=&"

	// append number of rows
	end += fmt.Sprintf("rows=%d", options.Rows)

	end += "&page=1&output=json"

	escaped := url.QueryEscape(end)

	// undo all the characters that shouldn't be escaped
	escaped = strings.Replace(escaped, "%26", "&", -1)
	escaped = strings.Replace(escaped, "%3D", "=", -1)
	escaped = strings.Replace(escaped, "%2B", "+", -1)

	return start + escaped
}

func AdvancedSearch(query Query, options Options) (searchResult, error) {
	// do HTTP GET request
	searchUrl := constructUrl(query, options)
	r, err := http.Get(searchUrl)
	if err != nil {
		return searchResult{}, err
	}
	defer r.Body.Close()

	// decode the response into JSON
	res := searchResult{}
	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return searchResult{}, err
	}

	return res, nil
}
