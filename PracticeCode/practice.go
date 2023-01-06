package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	counter := 0
	fmt.Println("goroutine times: ", runtime.NumGoroutine())
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("goroutine times: ", runtime.NumGoroutine())
	fmt.Println("count = ", counter)

}

func foo() {
	for i := 0; i < 3; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 3; i++ {
		fmt.Println("bar:", i)
	}
}

func grr() {
	for i := 0; i < 3; i++ {
		fmt.Println("grr:", i)
	}
	fmt.Println("done_2")
}
