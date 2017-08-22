package internetarchive

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type response struct {
	ResponseHeader header `json:"responseHeader"`
	Response       body   `json:"response"`
}
type header struct {
	Status float64                `json:"status"`
	QTime  float64                `json:"QTime"`
	Params map[string]interface{} `json:"params"`
}

type body struct {
	NumFound int   `json:"numFound"`
	Start    int   `json:"start"`
	Docs     []doc `json:"docs"`
}

type doc struct {
	Identifier string  `json:"identifier"`
	Downloads  float64 `json:"downloads"`
	Date       string  `json:"date"`
	AvgRating  string  `json:"avg_rating"`
}

func Response(url string) (response, error) {
	searchUrl := url

	r, err := http.Get(searchUrl)
	if err != nil {
		return response{}, err
	}
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return response{}, err
	}

	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return response{}, err
	}

	return resp, nil
}
