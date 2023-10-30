package main

import (
	"fmt"
	"sync"
)

func main() {
	finishAll()
	finishAllMutex()
	MoreReadMutex()
	MoreReadRWMutex()
}

func finishAll() {
	var counter int64
	var wg = sync.WaitGroup{}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}

func finishAllMutex() {
	var counter int64
	var m = sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			counter++
		}()
	}

	fmt.Println(counter)
}

func MoreReadMutex() {
	var counter int
	var m = sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func() {
			m.Lock()
			counter = 10
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			m.Unlock()
		}()
	}
}

func MoreReadRWMutex() {
	var counter int
	var m = sync.RWMutex{}
	for i := 0; i < 1000; i++ {
		go func() {
			m.Lock()
			counter = 10
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			_ = fmt.Sprint(counter)
			m.Unlock()
		}()
	}
}
