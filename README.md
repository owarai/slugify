# slugify


Package `slugify` generate a dash-connected slug from unicode string.

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/owarai/slugify?tab=doc)

## Example

```go
package main

import (
    "fmt"
    
    "github.com/owarai/slugify"
)

func main() {
    text := slugify.Format("hello, 你好，world! 世界！", true)
    fmt.Println(text) // Will print: "hello-你好-world-世界"
    
    someText := slugify.Format("hello, 你好，world! 世界！", false)
    fmt.Println(someText) // Will print: "hello-world"
}
```

## Installation
```sh
# under module-aware mode
go get github.com/owarai/slugify
```
