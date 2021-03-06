package url

import (
	"fmt"
	"testing"
)

var testValues = []struct {
	Input            string
	WantedSubdomains []string
	WantedDomain     string
	WantedTLD        string
	WantedCTLD       string
	WantedFullDomain string
	WantedPath       string
	WantedQueries    map[string][]string
	ShouldFail       bool
}{
	{
		Input:            "https://an.awesome.blog.boratanrikulu.dev.tr/blog/archlinux-install.html?q=a+lovely+query&z=another+query",
		WantedSubdomains: []string{"an", "awesome", "blog"},
		WantedDomain:     "boratanrikulu",
		WantedTLD:        "dev",
		WantedCTLD:       "tr",
		WantedFullDomain: "an.awesome.blog.boratanrikulu.dev.tr",
		WantedPath:       "/blog/archlinux-install.html",
		WantedQueries: map[string][]string{
			"q": []string{"a", "lovely", "query"},
			"z": []string{"another", "query"},
		},
		ShouldFail: false,
	},
	{
		Input:            "https://boratanrikulu.dev/dns-guvenlik-sorunlari",
		WantedSubdomains: []string{},
		WantedDomain:     "boratanrikulu",
		WantedTLD:        "dev",
		WantedCTLD:       "",
		WantedFullDomain: "boratanrikulu.dev",
		WantedPath:       "/dns-guvenlik-sorunlari",
		WantedQueries:    map[string][]string{},
		ShouldFail:       false,
	},
	{
		Input:            "https://boratanrikulu",
		WantedSubdomains: []string{},
		WantedDomain:     "",
		WantedTLD:        "",
		WantedCTLD:       "",
		WantedFullDomain: "",
		WantedPath:       "",
		WantedQueries:    map[string][]string{},
		ShouldFail:       true,
	},
	{
		Input:            "https://bora.fi",
		WantedSubdomains: []string{},
		WantedDomain:     "bora",
		WantedTLD:        "fi",
		WantedCTLD:       "",
		WantedFullDomain: "bora.fi",
		WantedPath:       "",
		WantedQueries:    map[string][]string{},
		ShouldFail:       false,
	},
	{
		Input:            "https://boratanrikulu.com.tr",
		WantedSubdomains: []string{},
		WantedDomain:     "boratanrikulu",
		WantedTLD:        "com",
		WantedCTLD:       "tr",
		WantedFullDomain: "boratanrikulu.com.tr",
		WantedPath:       "",
		WantedQueries:    map[string][]string{},
		ShouldFail:       false,
	},
	{
		Input:            "https://boratanrikulu.com.randomwrongctld",
		WantedSubdomains: []string{},
		WantedDomain:     "",
		WantedTLD:        "",
		WantedCTLD:       "",
		WantedFullDomain: "",
		WantedPath:       "",
		WantedQueries:    map[string][]string{},
		ShouldFail:       true,
	},
	{
		Input:            "https://boratanrikulu.randomwrongtld",
		WantedSubdomains: []string{},
		WantedDomain:     "",
		WantedTLD:        "",
		WantedCTLD:       "",
		WantedFullDomain: "",
		WantedPath:       "",
		WantedQueries:    map[string][]string{},
		ShouldFail:       true,
	},
	{
		Input:            "https://api.seo.do",
		WantedSubdomains: []string{"api"},
		WantedDomain:     "seo",
		WantedTLD:        "do",
		WantedCTLD:       "",
		WantedFullDomain: "api.seo.do",
		WantedPath:       "",
		WantedQueries:    map[string][]string{},
		ShouldFail:       false,
	},
	{
		Input:            "",
		WantedSubdomains: []string{},
		WantedDomain:     "",
		WantedTLD:        "",
		WantedCTLD:       "",
		WantedFullDomain: "",
		WantedPath:       "",
		WantedQueries:    map[string][]string{},
		ShouldFail:       true,
	},
}

func TestNewURL(t *testing.T) {
	for _, testValue := range testValues {
		fmt.Println(testValue.Input)

		u, err := NewURL(testValue.Input)
		if testValue.ShouldFail {
			if err != nil {
				continue
			} else {
				t.Fatalf("[%s] Error must be occurred, but did not", testValue.Input)
			}
		}

		if !testValue.ShouldFail && err != nil {
			t.Fatalf("Error occur: %s - %s", err, testValue.Input)
		}

		if !equalStringSlice(testValue.WantedSubdomains, u.Subdomains) {
			t.Fatalf("[%s] Subdomains are wrong: Wanted: \"%s\" - Got: \"%s\"", testValue.Input, testValue.WantedSubdomains, u.Subdomains)
		}

		if testValue.WantedDomain != u.Domain {
			t.Fatalf("[%s] Domain is wrong: Wanted: \"%s\" - Got: \"%s\"", testValue.Input, testValue.WantedDomain, u.Domain)
		}

		if testValue.WantedTLD != u.TLD {
			t.Fatalf("[%s] TLD is wrong: Wanted: \"%s\" - Got: \"%s\"", testValue.Input, testValue.WantedTLD, u.TLD)
		}

		if testValue.WantedCTLD != u.CTLD {
			t.Fatalf("[%s] CTLD is wrong: Wanted: \"%s\" - Got: \"%s\"", testValue.Input, testValue.WantedCTLD, u.CTLD)
		}

		if testValue.WantedPath != u.Path {
			t.Fatalf("[%s] Path is wrong: Wanted: \"%s\" - Got: \"%s\"", testValue.Input, testValue.WantedPath, u.Path)
		}

		if fmt.Sprint(testValue.WantedQueries) != fmt.Sprint(u.Queries) {
			t.Fatalf("[%s] Queries are wrong: Wanted: \"%s\" - Got: \"%s\"", testValue.Input, testValue.WantedQueries, u.Queries)
		}
	}
}

func TestIsLive(t *testing.T) {
	var testValues = []struct {
		Input string
		Want  bool
	}{
		{
			Input: "https://yagizdegirmenci.com",
			Want:  true,
		},
		{
			Input: "https://golang.org",
			Want:  true,
		},
		{
			Input: "https://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.com",
			Want:  false,
		},
		{
			Input: "https://randomsiteadi.com",
			Want:  false,
		},
	}

	for _, testValue := range testValues {
		u, err := NewURL(testValue.Input)
		if err != nil {
			t.Fatalf("Given raw url for the IsLive testing is not correct: %s", testValue.Input)
		}

		response := u.IsLive()
		if response != testValue.Want {
			t.Fatalf("[%s] Result from URL is wrong: Wanted: \"%t\" - Got: \"%t\"", testValue.Input, testValue.Want, response)
		}
	}
}

func TestIsRecorded(t *testing.T) {
	var testValues = []struct {
		Input string
		Want  bool
	}{
		{
			Input: "https://boratanrikulu.dev",
			Want:  true,
		},
		{
			Input: "https://boratanrikulu.dev/asdasd/asdasdas/asdasd",
			Want:  true,
		},
		{
			Input: "https://asdasd.asdasdas.asd.boratanrikulu.dev/asdasda/sdasd/asd/asd",
			Want:  false,
		},
		{
			Input: "https://api.seo.do/asdasda/sdasd/asd/asd",
			Want:  true,
		},
		{
			Input: "https://A8QAPDm68bjNc.org",
			Want:  false,
		},
	}

	for _, testValue := range testValues {
		u, err := NewURL(testValue.Input)
		if err != nil {
			t.Fatalf("Given raw url for the IsRecorded testing is not correct: %s", testValue.Input)
		}

		response := u.IsRecorded()
		if response != testValue.Want {
			t.Fatalf("[%s] Result from URL is wrong: Wanted: \"%t\" - Got: \"%t\"", testValue.Input, testValue.Want, response)
		}
	}
}

func equalStringSlice(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
