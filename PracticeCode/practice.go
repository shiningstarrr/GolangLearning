package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start")
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	var counter int64
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)

			//runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&counter))

			wg.Done()
			//fmt.Println("goroutines: ", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("count: ", counter)
	fmt.Println("goroutines: ", runtime.NumGoroutine())

}
