# Go Closure Compiler

Go bindings for Google's Closure Compiler tool.

## Requirements

[Closure Compiler](https://developers.google.com/closure/compiler) must be accessible from your system's $PATH environment variable.

Install the package

```
go get github.com/markgx/go-closure-compiler
```

## Usage

```
package main

import (
	"github.com/markgx/go-closure-compiler"
)

func main() {
	// optional options map
	options := make(map[string]string)
	options["--compilation_level"] = "WHITESPACE_ONLY"

	closure.Compile(&[]string{"backbone.js", "app.js"}, "app.min.js", options)
}
```

## License

Copyright 2013 Mark Guerra. Released under the [MIT License](http://www.opensource.org/licenses/MIT).
