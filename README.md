# hn

Go library for the [Hacker News API](https://github.com/HackerNews/API)

[![GoDoc](https://godoc.org/github.com/peterhellberg/hn?status.svg)](https://godoc.org/github.com/peterhellberg/hn)

## Installation

```bash
go get -u github.com/peterhellberg/hn
```

## Example usage

Showing the current top ten stories

```go
package main

import (
  "fmt"

  "github.com/peterhellberg/hn"
)

func main() {
  hn := hn.NewClient(nil)

  ids, err := hn.TopStories()
  if err != nil {
    panic(err)
  }

  for i, id := range ids[:10] {
    item, err := hn.Item(id)
    if err != nil {
      panic(err)
    }

    fmt.Println(i, "–", item.Title, "\n   ", item.URL, "\n")
  }
}
```
