# Opinionated environment variable parser for Golang

[![](https://img.shields.io/static/v1?label=Sponsor&message=%E2%9D%A4&logo=GitHub&color=%23fe8e86)](https://github.com/sponsors/akamensky) [![Go Reference](https://pkg.go.dev/badge/github.com/akamensky/envparse.svg)](https://pkg.go.dev/github.com/akamensky/envparse) [![Go Report Card](https://goreportcard.com/badge/github.com/akamensky/envparse)](https://goreportcard.com/report/github.com/akamensky/envparse) [![Coverage Status](https://coveralls.io/repos/github/akamensky/envparse/badge.svg?branch=master)](https://coveralls.io/github/akamensky/envparse?branch=master) [![Build Status](https://travis-ci.org/akamensky/envparse.svg?branch=master)](https://travis-ci.org/akamensky/envparse)

The goal of this project is to make it easy to parse environment variables into complex configuration structures.

Highlights:
* Best effort parsing -- even if errors happened, the correctly defined/provided values will be parsed
* Support for most types (except `complex`) and pointers
* Support for list and embedded structs (or lists of structs)
* Custom prefix (default `APP`)
* Hardcoded (intentionally) separator `_`
* All errors reported in 1 (see helper functions)

### Installation

To install and start using this package simply do:

```
$ go get -u -v github.com/akamensky/envparse
```

### Usage

Here is basic example of parsing environment variables into configuration struct:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/akamensky/envparse"
)

type config struct {
	Url      string `env:"name=someUrl,default=http://localhost:8080"`
	IntValue int    `env:"name=value_int,required"`
	ObjSlice []*Obj `env:"name=obj"`
}

type Obj struct {
	A string `env:"name=a"`
	B string `env:"name=b"`
}

func main() {
	c := &config{}
	fakeEnviron := []string{
		"APP_SOMEURL=https://github.com/akamensky/envparse",
		"APP_VALUEINT=456",
		"APP_OBJ_0_A=hello",
		"APP_OBJ_0_B=world",
		"APP_OBJ_1_A=it's a",
		"APP_OBJ_1_B=slice",
    }
	envparse.SetPrefix("APP")
	// Use os.Environ() to get list of environment variables
	if err := envparse.Parse(c, fakeEnviron); err != nil {
		panic(envparse.FullError(err))
	}
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		panic(err)
    }
	fmt.Println(string(b))
}
```

### Options

* `name=<str>` -- a name portion for this field in envvar
* `required` -- mark field as required, if no matching environment variable found will cause error to be returned (appended)
* `default=<str>` -- a way to define default value, cannot be provided for `required` fields. If both found, it will cause error to be added
