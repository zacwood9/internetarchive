# internetarchive
A Go wrapper for the archive.org Advanced Search API

# Usage
Create an AdvancedSearch object
```go
search := AdvancedSearch{
		Query{"GratefulDead", "1977-05-08 OR 77-05-08"},
		Options{[]string{"avg_rating", "date", "downloads", "identifier"}, "downloads desc", 2}}

```

Get the response from archive.org for the search

```go
response, err := search.Search()
if err != nil {
    panic(err.Error())
}
```

Access the results of the search
```go
items := response.Body.Items
for i := 0; i < len(items); i++ {
    fmt.Println(items[i])
}
```
Output: 
```
{gd77-05-08.sbd.hicks.4982.sbeok.shnf 615323 1977-05-08T00:00:00Z 4.79}
{gd1977-05-08.shure57.stevenson.29303.flac16 471270 1977-05-08T00:00:00Z 4.59}
```