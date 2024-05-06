package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Results struct {
	snis  []*SNI
	mutex sync.Mutex
}

func (res *Results) AddSNI(sni *SNI) {
	res.mutex.Lock()
	res.snis = append(res.snis, sni)
	res.mutex.Unlock()
}

func (res *Results) Show(config *Config) {
	var writer *os.File = nil

	if config.file != "" {
		var err error
		writer, err = os.Create(config.file)
		if err != nil {
			log.Fatal("Cannot open "+config.file, err)
		}
		defer writer.Close()
	}

	if config.debug {
		log.Printf("Sorting %d SNIs.\n", len(res.snis))
	}
	// sort SNI by tls version, by http version & by duration
	sort.Slice(res.snis, func(i, j int) bool {
		if res.snis[i].tls != res.snis[j].tls {
			return res.snis[i].tls > res.snis[j].tls
		}
		if res.snis[i].http != res.snis[j].http {
			return res.snis[i].http > res.snis[j].http
		}
		if res.snis[i].duration == res.snis[j].duration {
			return res.snis[i].domain < res.snis[j].domain
		}
		return res.snis[i].duration < res.snis[j].duration
	})

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "SNI", "TLS", "Proto", "duration", "Extra"})
	for i, sni := range res.snis {
		if config.best > 0 && (config.best-1) < i {
			break
		}
		var row []interface{}

		ok := sni.Test(config.h2only)
		if !config.debug && !ok {
			continue
		}

		if writer != nil && !config.debug {
			fmt.Fprintf(writer, "%s %s %s %v%s\n", sni.domain, sni.TLS(), sni.proto, sni.duration, sni.Extra())
		}

		row = []interface{}{
			i, sni.domain, sni.TLS(), sni.proto, sni.duration, sni.Extra(),
		}

		t.AppendRow(row)
	}
	t.Render()
}
