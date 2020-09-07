# godesu - 4chan read only API wrapper
Gochan desu is 4chan read only API wrapper written in Go.
It provides a simple API to get structs from HTTP API responses.

It has been written using standard lib only so there are no dependencies.
## Installation
`go get -u github.com/mtarnawa/godesu`
## Usage
See full example in `example/main.go`.

```go
package main

import "github.com/mtarnawa/godesu"

func main() {
    gochan := godesu.New()
    _, thread := gochan.Board("w").GetThread(1565459)
    
    images := thread.Images()
	for _, image := range images {
	  println(image.URL)
	  println(image.Filename)
	  println(image.Extension)
	  println(image.OriginalFilename)      
	}
}
```

## License
"Do-whatever-you-want-License 2.0

## Disclaimer
Not promoting 4chan, this is a helper utility for NSA.
