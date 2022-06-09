[![Go](https://github.com/gerbenjacobs/go-habbo/actions/workflows/go.yml/badge.svg)](https://github.com/gerbenjacobs/go-habbo/actions/workflows/go.yml)

# go-habbo

Library to deal with unofficial Habbo API.

## Installation

    go get github.com/go-habbo/habbo

## Usage

```go
parser := client.NewParser(http.DefaultClient)
api := client.NewHabboAPI(parser)

habbo, err := api.GetHabboByName(ctx, "com", "myHabboName")
if err != nil {
    // handle error
}
```
