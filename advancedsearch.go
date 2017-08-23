package internetarchive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type AdvancedSearch struct {
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

func (search AdvancedSearch) Url() string {
	start := "https://archive.org/advancedsearch.php?q="

	// format collection and query string
	end := fmt.Sprintf("collection:\"%s\" %s", search.Query.Collection, search.Query.SearchTerm)
	end += "&"

	// append all fields
	for i := 0; i < len(search.Options.Fields); i++ {
		end += fmt.Sprintf("fl[]=%s&", search.Options.Fields[i])
	}

	// append sort method
	end += fmt.Sprintf("sort[]=%s&", search.Options.SortBy)
	end += "sort[]=&sort[]=&"

	// append number of rows
	end += fmt.Sprintf("rows=%d", search.Options.Rows)

	end += "&page=1&output=json"

	escaped := url.QueryEscape(end)

	// undo all the characters that shouldn't be escaped
	escaped = strings.Replace(escaped, "%26", "&", -1)
	escaped = strings.Replace(escaped, "%3D", "=", -1)
	escaped = strings.Replace(escaped, "%2B", "+", -1)

	return start + escaped
}

func (search AdvancedSearch) Search() (result, error) {
	// do HTTP GET request
	r, err := http.Get(search.Url())
	if err != nil {
		return result{}, err
	}
	defer r.Body.Close()

	// read the HTTP response into a byte slice
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return result{}, err
	}

	// unmarshal the JSON body into res
	res := result{}
	if err := json.Unmarshal(data, &res); err != nil {
		return result{}, err
	}

	return res, nil
}
