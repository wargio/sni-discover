package main

import (
	"regexp"
	"strings"
)

var (
	SNI_RX = regexp.MustCompile(`(([a-z0-9-]{1,63}\.)?(xn--)?[a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,63}`)
	NS_RX  = regexp.MustCompile(`^ns(\d+)?\.`)
)

func trim(s string) string {
	return strings.TrimSpace(s)
}

func isSNI(s string) bool {
	return len(s) > 0 && SNI_RX.MatchString(s)
}

func removeNs(s string) string {
	return NS_RX.ReplaceAllString(strings.TrimSpace(s), "")
}
