package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

// MUTEXES ARE USED TO AVOID RACE CONDITIONS
func (p *post) inc(i int, wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += i
}

func main() {
	ps := post{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go ps.inc(i, &wg)
	}

	wg.Wait()

	fmt.Println(ps.views)
}
