package internetarchive

type searchResult struct {
	Header header `json:"responseHeader"`
	Body   body   `json:"response"`
}

type header struct {
	Status float64                `json:"status"`
	QTime  float64                `json:"QTime"`
	Params map[string]interface{} `json:"params"`
}

type body struct {
	NumFound int    `json:"numFound"`
	Start    int    `json:"start"`
	Items    []Item `json:"docs"`
}
