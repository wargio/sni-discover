package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

type SNI struct {
	domain      string
	duration    time.Duration
	tls         uint16
	http        int
	proto       string
	badCA       bool
	errored     error
	unreachable bool
}

func (i *SNI) Test(http2 bool) bool {
	if i.errored != nil || i.unreachable {
		return false
	} else if i.tls < tls.VersionTLS13 {
		return false
	} else if http2 && i.http < 20 {
		return false
	}
	return true
}

func (i *SNI) TLS() string {
	if i.errored != nil || i.unreachable {
		return ""
	}
	switch i.tls {
	case uint16(0):
		return "No TLS"
	case tls.VersionTLS10:
		return "TLSv1.0"
	case tls.VersionTLS11:
		return "TLSv1.1"
	case tls.VersionTLS12:
		return "TLSv1.2"
	case tls.VersionTLS13:
		return "TLSv1.3"
	default:
		return fmt.Sprintf("TLS(%x)", i.tls)
	}
}

func (i *SNI) Extra() string {
	if i.errored != nil {
		return i.errored.Error()
	} else if i.unreachable {
		return "Unreachable"
	} else if i.badCA {
		return "Bad CA"
	}
	return ""
}

func NewSNI(domain string, timeout time.Duration) *SNI {
	request, err := http.NewRequest("GET", "https://"+domain, nil)
	if err != nil {
		return &SNI{domain: domain, duration: timeout, errored: err}
	}

	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   timeout,
	}

	start := time.Now()
	response, err := client.Do(request)
	duration := time.Since(start)

	info := &SNI{domain: domain, duration: duration}

	// When http status code is not 200 OK
	if response == nil {
		info.duration = timeout
		info.unreachable = true
		return info
	}

	defer response.Body.Close()
	if response.TLS == nil {
		// no tls
		return info
	}

	info.tls = response.TLS.Version
	info.proto = response.Proto
	info.http = (response.ProtoMajor * 10) + response.ProtoMinor
	info.badCA = len(response.TLS.VerifiedChains) == 0
	return info
}
