# ZEO.ORG - URL

A tiny package that you extract your URLs by validating them.

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/zeoagency/url"
)

func main() {
	u, err := url.NewURL("https://boratanrikulu.dev/dns-guvenlik-sorunlari/")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(u.Domain) // boratanrikulu
	fmt.Println(u.Path)   // /dns-guvenlik-sorunlari
}
```
