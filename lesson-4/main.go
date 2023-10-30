package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var UserValue int

func main() {
	w := NewWorker()
	ch := make(chan os.Signal, 1)
	w.Start()
	defer w.Stop()
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	time.Sleep(2 * time.Second)
	_ = <-ch

}

type Worker struct {
	done chan struct{}
}

func NewWorker() *Worker {
	return &Worker{
		done: make(chan struct{}),
	}
}

func (w *Worker) do() {
	if UserValue != 1000 {
		UserValue++
		fmt.Println(UserValue)
	}
}

func (w *Worker) Start() {
	ticker := time.NewTicker(time.Microsecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.do()
			case <-w.done:
				ticker.Stop()
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.done)
}
