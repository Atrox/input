# Input

[![Build Status](https://img.shields.io/travis/Atrox/input.svg?style=flat-square)](https://travis-ci.org/Atrox/input)
[![Coverage Status](https://img.shields.io/coveralls/Atrox/input.svg?style=flat-square)](https://coveralls.io/r/Atrox/input)
[![Go Report Card](https://goreportcard.com/badge/github.com/atrox/input?style=flat-square)](https://goreportcard.com/report/github.com/atrox/input)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/Atrox/input)

> Simple, easy to use input handler for the CLI

## Installation

```sh
go get github.com/atrox/input
# or with dep
dep ensure -add github.com/atrox/input
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/atrox/input"
)

func main() {
	// Input (can be empty)
	someInput := input.Prompt("How was your day").(string)

	// Required Input
	someInput = input.Prompt("What is your favourite tv-show", input.RequiredValidator).(string)

	// Boolean Input (y/n)
	boolInput := input.Prompt("Are you sure (y/n)", input.RequiredValidator, input.BooleanValidator).(bool)

	// Boolean Input with default (Y/n)
	var inputResult bool
	switch boolInput2 := input.Prompt("Are you sure (Y/n)", input.BooleanValidator).(type) {
	case bool:
		// user set it with yes or no
		inputResult = boolInput2
	default:
		// user just pressed enter - set default
		inputResult = true
	}

	// File Input
	file := input.Prompt("Location of the config file?", input.RequiredValidator, input.FileValidator).(string)
	fmt.Println("File location:", file)
}

```

## Validators

List of built-in validators:

- RequiredValidator: ensures the input is not empty
- PathValidator: ensures the input is valid looking path
- DirectoryValidator: ensures the input is a valid and **existing** directory
- FileValidator: ensures the input is a valid and **existing** file
- IntegerValidator: ensures and converts the input to a integer with [strconv.Atoi](https://golang.org/pkg/strconv/#Atoi)
- BooleanValidator: ensures and converts the input to a boolean
    - **true**: 1, t, true, y, yes
    - **false**: 0, f, false, n, no

Create your own - you just need to create a `input.ValidatorFunction`:

```go
func ContainsSomethingValidator(input string) (interface{}, error) {
    if !strings.Contains("something") {
        return nil, fmt.Errorf("Input %s does not contain 'something'", input)
    }

    return input, nil
}
```

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/atrox/input/issues)
- Fix bugs and [submit pull requests](https://github.com/atrox/input/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features
