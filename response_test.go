package internetarchive

import "testing"

func TestResponse(t *testing.T) {
	search := AdvancedSearch{
		Query{"GratefulDead", "1977-05-08 OR 77-05-08"},
		Params{[]string{"avg_rating", "date", "downloads", "identifier"}, "downloads desc", 1}}

	_, err := Response(search.Url())
	if err != nil {
		t.Fatal(err)
	}

}
