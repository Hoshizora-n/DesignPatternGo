package objectpool

import (
	"fmt"
	"time"
)

type Worker struct {
	id int
}

func NewWorker(id int) *Worker {
	return &Worker{id}
}

func (w *Worker) Work() {
	fmt.Printf("Worker %d is working\n", w.id)
	time.Sleep(2 * time.Second)
}

type ObjectPool struct {
	workers chan *Worker
}

func NewObjectPool(capacity int) *ObjectPool {
	pool := &ObjectPool{
		workers: make(chan *Worker, capacity),
	}

	for i := 1; i <= capacity; i++ {
		pool.workers <- NewWorker(i)
	}

	return pool
}

func (p *ObjectPool) GetWorker() *Worker {
	return <-p.workers
}

func (p *ObjectPool) ReleaseWorker(worker *Worker) {
	p.workers <- worker
}

func Run() {
	pool := NewObjectPool(4)

	for i := 0; i < 10; i++ {
		worker := pool.GetWorker()
		go func(w *Worker) {
			w.Work()
			pool.ReleaseWorker(w)
		}(worker)
	}

	time.Sleep(1 * time.Second)
}
