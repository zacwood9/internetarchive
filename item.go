package internetarchive

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Item struct {
	Identifier string   `json:"identifier"`
	Downloads  float64  `json:"downloads"`
	Date       string   `json:"date"`
	AvgRating  string   `json:"avg_rating"`
	Creator    string   `json:"creator"`
	Title      string   `json:"title"`
	Metadata   metadata `json:"-"`
}

type metadata struct {
	Created    float64                  `json:"created"`
	Dir        string                   `json:"dir"`
	Files      []map[string]interface{} `json:"files"`
	FilesCount float64                  `json:"files_count"`
	ItemSize   float64                  `json:"item_size"`
	Reviews    []map[string]interface{} `json:"reviews"`
	CustomData map[string]interface{}   `json:"metadata"`
}

func (item *Item) GetMetadata() error {
	requestUrl := "http://archive.org/metadata/" + item.Identifier

	// GET request to Archive Metadata API
	r, err := http.Get(requestUrl)
	if err != nil {
		return err
	}

	// read response
	data, err := ioutil.ReadAll(r.Body)
	check(err)
	// unmarshal response
	if err := json.Unmarshal(data, &item.Metadata); err != nil {
		return err
	}

	return nil
}
