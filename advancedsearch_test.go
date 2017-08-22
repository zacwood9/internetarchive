package internetarchive

import (
	"testing"
)

func TestAdvancedSearch_Url(t *testing.T) {
	search := AdvancedSearch{
		Query{"GratefulDead", "1977-05-08 OR 77-05-08"},
		Params{[]string{"avg_rating", "date", "downloads", "identifier"}, "downloads desc", 50}}

	expected := "https://archive.org/advancedsearch.php?q=collection%3A%22GratefulDead%22+1977-05-08+OR+77-05-08&fl%5B%5D=avg_rating&fl%5B%5D=date&fl%5B%5D=downloads&fl%5B%5D=identifier&sort%5B%5D=downloads+desc&sort%5B%5D=&sort%5B%5D=&rows=50&page=1&output=json"
	received := search.Url()
	if received != expected {
		t.Fatal("Expected and generated URLs do not match")
	}
}
