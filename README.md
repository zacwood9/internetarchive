# internetarchive
A Go wrapper for the archive.org Advanced Search API

# Usage
Create an AdvancedSearch object
```go
search := internetarchive.AdvancedSearch{
		Query{"GratefulDead", "1977-05-08 OR 77-05-08"},
		Params{[]string{"avg_rating", "date", "downloads", "identifier"}, "downloads desc", 2}}
```

Get the response from archive.org for the search

```go
response, err := internetarchive.Response(search.Url())
if err != nil {
    panic(err)
}
```

Access the results of the search
```go
 docs := response.Response.Docs
	for i := 0; i < len(docs); i++ {
		fmt.Println(docs[i])
	}
```
Output: 
```
{gd77-05-08.sbd.hicks.4982.sbeok.shnf 615323 1977-05-08T00:00:00Z 4.79}
{gd1977-05-08.shure57.stevenson.29303.flac16 471270 1977-05-08T00:00:00Z 4.59}
```