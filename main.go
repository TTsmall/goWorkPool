package main

import (
	"fmt"
	"goWorkPool/gopool"
	"time"
)

type addRepChan struct {
	cmd int
	dec string
}

func (this *addRepChan) Execute() {
	for count := 0; count < 10000; count++ {
		if count > 1000000000 {
			break
		}
	}
}

func main() {
	var routines int = 10
	var items int = 100000
	fmt.Println("1 time=%s", time.Now().Format("2006-01-02 15:04:05.000"))
	workPools := gopool.NewPoolMultiParam(routines, items)
	for count := 0; count < 1000000; count++ {
		workPools.Add(&addRepChan{cmd: count})
	}
	workPools.Stop()
	fmt.Println("2 time=%s", time.Now().Format("2006-01-02 15:04:05.000"))
	for count := 0; count < 1000000; count++ {
		addChan := &addRepChan{cmd: count}
		addChan.Execute()
	}
	fmt.Println("3 time=%s", time.Now().Format("2006-01-02 15:04:05.000"))
}
