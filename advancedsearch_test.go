package internetarchive

import (
	"testing"
)

func TestConstructUrl(t *testing.T) {
	query := Query{"GratefulDead", "1977-05-08 OR 77-05-08"}
	options := 	Options{[]string{"avg_rating", "date", "downloads", "identifier"}, "downloads desc", 50}

	expected := "https://archive.org/advancedsearch.php?q=collection%3A%22GratefulDead%22+1977-05-08+OR+77-05-08&fl%5B%5D=avg_rating&fl%5B%5D=date&fl%5B%5D=downloads&fl%5B%5D=identifier&sort%5B%5D=downloads+desc&sort%5B%5D=&sort%5B%5D=&rows=50&page=1&output=json"
	received := constructUrl(query, options)
	if received != expected {
		t.Fatal("Expected and generated URLs do not match")
	}
}

func TestAdvancedSearch(t *testing.T) {
	query := Query{"GratefulDead", "1977-05-08 OR 77-05-08"}
	options := Options{[]string{"avg_rating", "date", "downloads", "identifier"}, "downloads desc", 1}

	_, err := AdvancedSearch(query, options)
	if err != nil {
		t.Fatal("Search failed", err)
	}
}
