package internetarchive

import (
	"fmt"
	"net/url"
	"strings"
)

type AdvancedSearch struct {
	Query  Query
	Params Params
}

func (search AdvancedSearch) Url() string {
	start := "https://archive.org/advancedsearch.php?q="
	end := search.Query.string()
	end += "&"
	for i := 0; i < len(search.Params.Fields); i++ {
		end += fmt.Sprintf("fl[]=%s&", search.Params.Fields[i])
	}
	end += fmt.Sprintf("sort[]=%s&", search.Params.SortBy)
	end += "sort[]=&sort[]=&"
	end += fmt.Sprintf("rows=%d", search.Params.Rows)
	end += "&page=1&output=json"

	escaped := url.QueryEscape(end)
	escaped = strings.Replace(escaped, "%26", "&", -1)
	escaped = strings.Replace(escaped, "%3D", "=", -1)
	escaped = strings.Replace(escaped, "%2B", "+", -1)
	return start + escaped
}

type Query struct {
	Collection string
	SearchTerm string
}

type Params struct {
	Fields []string
	SortBy string
	Rows   int
}

func (q Query) string() string {
	return fmt.Sprintf("collection:\"%s\" %s", q.Collection, q.SearchTerm)
}
