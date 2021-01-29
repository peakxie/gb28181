package gr

import "sync"

var wg sync.WaitGroup

func StartGoRoutine(f func()) {
	wg.Add(1)
	go func() {
		f()
		wg.Done()
	}()
}

func Wait() {
	wg.Wait()
}

func Add(delta int) {
	wg.Add(delta)
}

func Done() {
	wg.Done()
}
