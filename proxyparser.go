package proxyparser

import (
	"net/url"
	"strings"

	"golang.org/x/exp/slices"
)

type ProxyData struct {
	Scheme   string
	Hostname string
	Port     string

	// user:password format
	UserString string
}

func (d ProxyData) String() string {
	if d.Hostname == "" || d.Port == "" {
		return ""
	}

	var ret string = d.Scheme + "://"
	if d.UserString != "" {
		ret += d.UserString + "@"
	}

	ret += d.Hostname + ":" + d.Port
	return ret
}

var commonSchemes []string = []string{
	"http",
	"https",
	"socks4",
	"socks5",
}

// Allowed Formats: {scheme://user:pass@host:port}, {scheme://host:port}
//
// The scheme can be NOT provided. "http" will be provided by default.
//
// After the scheme there must ALWAYS be either ":" or "://"
//
// Returns {ProxyData, SuccessParsing}
func GetProxyData(rawProxy string) (ProxyData, bool) {
	var data ProxyData = ProxyData{}

	// if there's provied a shit scheme, well.. remove it
	if strings.Contains(rawProxy, "://") {
		if split := strings.Split(rawProxy, "://"); len(split) > 0 {
			if !slices.Contains(commonSchemes, strings.ToLower(split[0])) {
				rawProxy = rawProxy[len(split[0]+"://"):]
			}
		}
	}

	var hasScheme bool = false
	for _, scheme := range commonSchemes {
		if strings.HasPrefix(strings.ToLower(rawProxy), scheme) {
			// now we make sure we have the scheme formatted right
			var afterScheme string = rawProxy[len(scheme):]
			afterScheme = strings.TrimPrefix(afterScheme, ":")
			afterScheme = strings.TrimPrefix(afterScheme, "//")
			rawProxy = scheme + "://" + afterScheme

			hasScheme = true
			break
		}
	}

	if !hasScheme {
		rawProxy = "http://" + rawProxy
	}

	var url, err = url.Parse(rawProxy)
	if err != nil {
		// fmt.Println(err)
		return data, false
	}

	// should NEVER happen, because we already add a scheme to rawProxy up
	if url.Scheme == "" {
		return data, false
	}

	// ayo
	if url.Hostname() == "" {
		return data, false
	}

	data.Scheme = url.Scheme
	data.Hostname = url.Hostname()
	data.Port = url.Port()
	if url.User != nil {
		data.UserString = url.User.String()
	}

	return data, true
}
