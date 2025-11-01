package util

import "strings"

func DomainKey(host string) string {
	host = strings.ToLower(host)
	if strings.HasPrefix(host, "www.") {
		return strings.TrimPrefix(host, "www.")
	}
	return host
}
