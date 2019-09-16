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
	var items int = 1000
	fmt.Println("1 time=", time.Now().Format("2006-01-02 15:04:05.000"))
	workPools := gopool.NewPoolMultiParam(routines, items)
	for count := 0; count < 1000000; count++ {
		workPools.Add(&addRepChan{cmd: count})
	}
	workPools.Stop()
	fmt.Println("2 time=", time.Now().Format("2006-01-02 15:04:05.000"))
	for count := 0; count < 1000000; count++ {
		addChan := &addRepChan{cmd: count}
		addChan.Execute()
	}
	fmt.Println("3 time=", time.Now().Format("2006-01-02 15:04:05.000"))
}


1 time= 2019-09-16 12:29:19.825
2 time= 2019-09-16 12:29:24.970
3 time= 2019-09-16 12:29:44.603
