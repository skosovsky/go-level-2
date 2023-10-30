package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

func main() {
	//trace.Start(os.Stderr) // go run main.go 2>trace.out и go tool trace trace.out
	//defer trace.Stop()

	safeMutexAdd()
	fmt.Println(counter)

	safeMutexDif()
	fmt.Println(counter)
}

func safeMutexAdd() {
	var m = sync.Mutex{}
	var wg = sync.WaitGroup{}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
}

// Закомментируя здесь Mutex - получим гонку
func safeMutexDif() {
	var m = sync.Mutex{}
	var wg = sync.WaitGroup{}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			runtime.Gosched() // Явный вызов планироващика
			counter--
		}()
	}
	wg.Wait()
}
