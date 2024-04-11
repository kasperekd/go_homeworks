package main

import (
	"fmt"
	"sync"
)

type TurnstileCounter struct {
	count int
	mutex sync.Mutex
}

func (tc *TurnstileCounter) incrementCounter() {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()
	tc.count++
}

func main() {
	turnstile := TurnstileCounter{}

	passThroughTurnstile := func(wg *sync.WaitGroup, i int) {
		defer wg.Done()
		turnstile.incrementCounter()
		fmt.Printf("Человек прошел через турникет %d\n", i)
	}

	var wg sync.WaitGroup
	wg.Add(20000)

	for i := 0; i < 20000; i++ {
		go passThroughTurnstile(&wg, i)
	}

	wg.Wait()

	fmt.Printf("Всего людей прошло через турникет: %d\n", turnstile.count)
}
