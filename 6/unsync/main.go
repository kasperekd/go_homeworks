package main

import (
	"fmt"
	"sync"
)

type TurnstileCounter struct {
	count int
}

func (tc *TurnstileCounter) incrementCounter() {
	tc.count++
}

func main() {
	turnstile := TurnstileCounter{}

	passThroughTurnstile := func(i int) {
		turnstile.incrementCounter()
		fmt.Printf("Человек прошел через турникет %d\n", i)
	}

	var wg sync.WaitGroup
	wg.Add(20000)

	for i := 0; i < 20000; i++ {
		go func() {
			passThroughTurnstile(i)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Фактическое количество прошедших людей может быть не корректным: %d\n", turnstile.count)
}
