package goroutine

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
//考察点 ： go 关键字的使用、协程的并发执行。

func GoroutinePrintNum() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			// 输出奇数 odd
			fmt.Println("输出的奇数：", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			// 输出偶数 even
			fmt.Println("输出的偶数：", i)
		}
	}()
	wg.Wait()
	fmt.Println("finish progress")
}

//题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
//考察点 ：协程原理、并发任务调度。

func MyTask(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	ranD := rand.IntN(1000)
	time.Sleep(time.Duration(ranD) * time.Millisecond)
	fmt.Println("执行第", i, "个任务 ,用时：", ranD, "ms.")
}

func ProgressTask(tasks []func(i int, wg *sync.WaitGroup)) {
	wg := &sync.WaitGroup{}
	wg.Add(len(tasks))
	for i, task := range tasks {
		go task(i, wg)
	}
	wg.Wait()
}

func DoTask() {
	tasks := make([]func(int, *sync.WaitGroup), 10)
	for i := range tasks {
		tasks[i] = MyTask
	}
	ProgressTask(tasks)
}
