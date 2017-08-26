# internetarchive
A Go wrapper for the archive.org Advanced Search API

# Usage
Construct Query and Options
```go
query := internetarchive.Query{"GratefulDead", "1977-05-08 OR 77-05-08"}
options := internetarchive.Options{[]string{"avg_rating", "date", "downloads", "identifier"}, "downloads desc", 2}
```

Perform a search using query and options
```go
response, err := internetarchive.AdvancedSearch(query, options)
if err != nil {
    panic(err.Error())
}
```

Access the results of the search
```go
items := response.Body.Items
for _, item := range items {
    fmt.Println(item)
}
```
Output: 
```
{gd77-05-08.sbd.hicks.4982.sbeok.shnf 615323 1977-05-08T00:00:00Z 4.79}
{gd1977-05-08.shure57.stevenson.29303.flac16 471270 1977-05-08T00:00:00Z 4.59}
```