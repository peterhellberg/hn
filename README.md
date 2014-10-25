# hn

Go library for the [Hacker News API](https://github.com/HackerNews/API)

[![GoDoc](https://godoc.org/github.com/peterhellberg/hn?status.svg)](https://godoc.org/github.com/peterhellberg/hn)
[![Build Status](https://travis-ci.org/peterhellberg/hn.svg?branch=master)](https://travis-ci.org/peterhellberg/hn)

## Installation

```bash
go get -u github.com/peterhellberg/hn
```

## Services

The client currently delegates to implementations of three interfaces:
[ItemsService](https://godoc.org/github.com/peterhellberg/hn#ItemsService),
[LiveService](https://godoc.org/github.com/peterhellberg/hn#LiveService) and
[UsersService](https://godoc.org/github.com/peterhellberg/hn#UsersService).

## Example usage

Showing the current top ten stories

```go
package main

import (
  "fmt"

  "github.com/peterhellberg/hn"
)

func main() {
  // NewClient returns a new Hacker News API client.
  // If httpClient is nil, http.DefaultClient is used.
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

Showing the current top ten stories using goroutines, a channel and a wait group

```go
package main

import (
  "fmt"
  "net/http"
  "sync"
  "time"

  "github.com/peterhellberg/hn"
)

type indexItem struct {
  Index int
  Item  *hn.Item
}

var (
  items    = map[int]*hn.Item{}
  messages = make(chan indexItem)
)

func main() {
  hn := hn.NewClient(&http.Client{
    Timeout: time.Duration(5 * time.Second),
  })

  ids, err := hn.TopStories()
  if err != nil {
    panic(err)
  }

  go func() {
    for i := range messages {
      items[i.Index] = i.Item
    }
  }()

  var wg sync.WaitGroup

  for i, id := range ids[:10] {
    wg.Add(1)
    go func(i, id int) {
      defer wg.Done()

      item, err := hn.Item(id)
      if err != nil {
        panic(err)
      }

      messages <- indexItem{i, item}
    }(i, id)
  }

  wg.Wait()

  for i := 0; i < 10; i++ {
    fmt.Println(i, "–", items[i].Title, "\n   ", items[i].URL, "\n")
  }
}
```

Showing information about a given user (first argument)

```go
package main

import (
  "fmt"
  "os"

  "github.com/peterhellberg/hn"
)

func main() {
  if u, err := hn.NewClient(nil).User(os.Args[1]); err == nil {
    fmt.Println("ID:   ", u.ID)
    fmt.Println("About:", u.About)
    fmt.Println("Karma:", u.Karma)
  }
}
```

## License

> *The MIT License (MIT)*
>
> Copyright (c) 2014 [Peter Hellberg](http://c7.se/)
>
> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included in all
> copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
> SOFTWARE.
