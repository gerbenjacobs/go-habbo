[![Go](https://github.com/gerbenjacobs/go-habbo/actions/workflows/go.yml/badge.svg)](https://github.com/gerbenjacobs/go-habbo/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/gerbenjacobs/go-habbo.svg)](https://pkg.go.dev/github.com/gerbenjacobs/go-habbo)

# go-habbo

Library to deal with unofficial Habbo API.

## Installation

    go get github.com/go-habbo/habbo

## Usage

Create a parser and a Habbo API client. Then you can fetch Habbos by name or by HHID.

```go
parser := client.NewParser(http.DefaultClient)
api := client.NewHabboAPI(parser)

habbo, err := api.GetHabboByName(ctx, "com", "myHabboName")
if err != nil {
    // handle error
}
```

When you have a Habbo object, you can use it to get more information about the Habbo.

Make sure to use the `hhid` (Habbo Unique ID).

```go
profile, err := api.GetProfile(ctx, "com", "hhus-123456789")
if err != nil {
    // handle error
}
```


## Habbo CLI

This projects also comes with a small command-line interface to fetch Habbos information as JSON.

```bash
$ go build -o go-habbo cmd/cli/main.go 
$ ./go-habbo <hotel> <habboName>
```

You can also run it straight from the source.

```bash
$ go run cmd/cli/main.go <hotel> <habboName>
```

The JSON is indented for readability and can be piped to a file or `jq` for further processing.

```bash
$ go run cmd/cli/main.go <hotel> <habboName> | jq .motto
```