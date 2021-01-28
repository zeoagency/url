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
	u, err := url.NewURL("https://an.awesome.blog.boratanrikulu.dev.tr/blog/archlinux-install.html?q=a+lovely+query&z=another+query")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(u.Subdomains)     // ["an", "awesome", "blog"]
	fmt.Println(u.Domain)         // "boratanrikulu"
	fmt.Println(u.TLD)            // "dev"
	fmt.Println(u.CTLD)           // "tr"
	fmt.Println(u.FullDomain)     // "an.awesome.blog.boratanrikulu.dev.tr"
	fmt.Println(u.Path)           // "/blog/archlinux-install.html"
	fmt.Println(u.Queries)        // map[q:["a", "lovely", "query"], z:["another", "query"]
	fmt.Println(u.IsLive())       // false
	fmt.Println(u.IsRecorded())   // false
}
```
