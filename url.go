package url

import (
	"errors"
	nethttp "net/http"
	neturl "net/url"
	"strings"
	"time"
)

// URL is a struct that you can extract each element of an URL.
//
// It works like this:
//
// Given URL:    "https://an.awesome.blog.boratanrikulu.dev.tr/blog/archlinux-install.html?q=a+lovely+query&z=another+query"
// Result:
//   SUB_DOMAINS: an.awesome.blog
//   DOMAIN:      boratanrikulu"
//   TLD:         dev
//   C-TLD:       tr
//   Path:        /blog/archlinux-install.html
//   Queries:     q:a+lovely+query, z:another+query
//
// Example Usage:
//
// u := NewURL("https://an.awesome.blog.boratanrikulu.dev.tr/blog/archlinux-install.html?q=a+lovely+query&z=an+other+query")
// fmt.Println(u.TLD // "dev"
// fmt.Println(u.CTLD) // "tr"
// fmt.Println(u.Queries) // "[q:[a lovely query] z:[another query]]"
type URL struct {
	Rawurl     string
	Subdomains []string
	Domain     string
	TLD        string
	CTLD       string
	Path       string
	Queries    map[string][]string
}

// NewURL returns a new URL by validating it.
func NewURL(rawurl string) (*URL, error) {
	u, err := neturl.Parse(rawurl)
	if err != nil || u.Scheme == "" {
		return nil, errors.New("That's not a valid URL.")
	}

	parts := strings.Split(u.Hostname(), ".")

	// Check if it has valid country tld at the last part.
	// Increase the tldCount if it does.
	tldCount := 1
	if _, c := stringSliceContains(countryTopLevelDomains, parts[len(parts)-1]); c {
		tldCount++
	}
	if tldCount >= len(parts) {
		return nil, errors.New("That's not a valid URL.")
	}

	// TLD
	tld := parts[len(parts)-tldCount]

	// Country TLD
	var ctld string
	if tldCount == 2 {
		ctld = parts[len(parts)-1]
	}

	// Domain
	domain := parts[len(parts)-tldCount-1]
	subDomains := parts[:len(parts)-tldCount-1]

	return &URL{
		Rawurl:     rawurl,
		Subdomains: subDomains,
		Domain:     domain,
		TLD:        tld,
		CTLD:       ctld,
		Path:       u.EscapedPath(),
		Queries:    u.Query(),
	}, nil
}

// stringSliceContains tells whether a contains x.
func stringSliceContains(a []string, x string) (int, bool) {
	for i, n := range a {
		if x == n {
			return i, true
		}
	}
	return -1, false
}

// isLive returns whether is URL.Rawurl is live or not.
func (u *URL) isLive() bool {
	// Set timeout.
	client := nethttp.Client{
		Timeout: 2 * time.Second,
	}
	_, err := client.Get(u.Rawurl)

	if err != nil {
		return false
	}

	return true
}

// countryTopLevelDomains includes all country-code-top-level-domains.
// source: https://en.wikipedia.org/wiki/List_of_Internet_top-level_domains#Country_code_top-level_domains
var countryTopLevelDomains = []string{
	"ac", "ad", "ae", "af", "ag", "ai", "al", "am", "ao", "aq", "ar", "as", "at", "au", "aw", "ax", "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo", "bq", "br", "bs", "bt", "bw", "by", "bz", "ca", "cc", "cd", "cf", "cg", "ch", "ci", "ck", "cl", "cm", "cn", "co", "cr", "cu", "cv", "cw", "cx", "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "eh", "er", "es", "et", "eu", "fi", "fj", "fk", "fm", "fo", "fr", "ga", "gd", "ge", "gf", "gg", "gh", "gi", "gl", "gm", "gn", "gp", "gq", "gr", "gs", "gt", "gu", "gw", "gy", "hk", "hm", "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "io", "iq", "ir", "is", "it", "je", "jm", "jo", "jp", "ke", "kg", "kh", "ki", "km", "kn", "kp", "kr", "kw", "ky", "kz", "la", "lb", "lc", "li", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me", "mg", "mh", "mk", "ml", "mm", "mn", "mo", "mp", "mq", "mr", "ms", "mt", "mu", "mv", "mw", "mx", "my", "mz", "na", "nc", "ne", "nf", "ng", "ni", "nl", "no", "np", "nr", "nu", "nz", "om", "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pm", "pn", "pr", "ps", "pt", "pw", "py", "qa", "re", "ro", "rs", "ru", "rw", "sa", "sb", "sc", "sd", "se", "sg", "sh", "si", "sk", "sl", "sm", "sn", "so", "sr", "ss", "st", "su", "sv", "sx", "sy", "sz", "tc", "td", "tf", "tg", "th", "tj", "tk", "tl", "tm", "tn", "to", "tr", "tt", "tv", "tw", "tz", "ua", "ug", "uk", "us", "uy", "uz", "va", "vc", "ve", "vg", "vi", "vn", "vu", "wf", "ws", "ye", "yt", "za", "zm", "zw",
}
