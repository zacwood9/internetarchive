package internetarchive

type result struct {
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

type Item struct {
	Identifier string  `json:"identifier"`
	Downloads  float64 `json:"downloads"`
	Date       string  `json:"date"`
	AvgRating  string  `json:"avg_rating"`
}
