package gopool

import (
	"sync"
)

type Item interface {
	Execute()
}

type Worker func()

type Pool struct {
	items chan Item
	once  sync.Once
	done  chan struct{}
	wg    sync.WaitGroup
}

func (this Worker) Execute() {
	this()
}

func NewPool(count int) *Pool {
	pool := &Pool{
		items: make(chan Item, count*10),
		done:  make(chan struct{}),
	}
	for i := 0; i < count; i++ {
		pool.wg.Add(1)
		go pool.run()
	}
	return pool
}

func NewPoolMultiParam(count, size int) *Pool {
	pool := &Pool{
		items: make(chan Item, size),
		done:  make(chan struct{}),
	}
	for i := 0; i < count; i++ {
		pool.wg.Add(1)
		go pool.run()
	}
	return pool
}

func (this *Pool) Add(item Item) {
	this.items <- item
}

func (this *Pool) run() {
	defer this.wg.Done()
	for {
		select {
		case item := <-this.items:
			item.Execute()
		case <-this.done:
			for {
				select {
				case item := <-this.items:
					item.Execute()
				default:
					return
				}
			}
		}
	}
}

func (this *Pool) Stop() {
	this.once.Do(func() {
		close(this.done)
	})
	this.wg.Wait()
}
