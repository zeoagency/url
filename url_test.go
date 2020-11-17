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
	WantedPath       string
	WantedQueries    map[string][]string
	ShouldFail       bool
}{
	{
		Input:            "https://boratanrikulu.dev/dns-guvenlik-sorunlari",
		WantedSubdomains: []string{},
		WantedDomain:     "boratanrikulu",
		WantedTLD:        "dev",
		WantedCTLD:       "",
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
		WantedPath:       "",
		WantedQueries:    map[string][]string{},
		ShouldFail:       true,
	},
	{
		Input:            "https://boratanrikulu.tr",
		WantedSubdomains: []string{},
		WantedDomain:     "",
		WantedTLD:        "",
		WantedCTLD:       "",
		WantedPath:       "",
		WantedQueries:    map[string][]string{},
		ShouldFail:       true,
	},
	{
		Input:            "https://boratanrikulu.com.tr",
		WantedSubdomains: []string{},
		WantedDomain:     "boratanrikulu",
		WantedTLD:        "com",
		WantedCTLD:       "tr",
		WantedPath:       "",
		WantedQueries:    map[string][]string{},
		ShouldFail:       false,
	},
}

func TestExtractURL(t *testing.T) {
	for _, testValue := range testValues {
		fmt.Println(testValue.Input)

		u, err := NewURL(testValue.Input)
		if testValue.ShouldFail {
			if err != nil {
				continue
			} else {
				t.Fatalf("[%s] Error must be occured, but did not", testValue.Input)
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