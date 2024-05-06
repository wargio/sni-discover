package main

import (
	"log"
	"sync"
	"time"
)

type Resolver struct {
	queue Queue
}

func sniResolver(rsv *Resolver, res *Results, wg *sync.WaitGroup, timeout time.Duration) {
	for {
		s := rsv.queue.Pop()
		if s == nil {
			break
		}
		sni := NewSNI(s.(string), timeout)
		res.AddSNI(sni)
	}
	wg.Done()
}

func (rsv *Resolver) AddSNI(sni string) {
	rsv.queue.Push(sni)
}

func (rsv *Resolver) Run(config *Config, res *Results) {
	var wg sync.WaitGroup
	timeout := time.Duration(config.timeout) * time.Second

	nSnis := rsv.queue.Size()
	if nSnis < 1 {
		return
	}

	if config.debug {
		log.Printf("Resolving %d SNIs.\n", nSnis)
	}

	routines := config.routines
	if routines > nSnis {
		routines = nSnis
	}

	for i := 0; i < routines; i++ {
		wg.Add(1)
		go sniResolver(rsv, res, &wg, timeout)
	}

	if config.debug {
		log.Println("Waiting for resolvers.")
	}

	wg.Wait()
}
