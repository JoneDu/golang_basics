package lock

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//✅锁机制
//题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ： sync.Mutex 的使用、并发数据安全。

type Counter struct {
	Num  int
	Lock sync.Mutex
}

func (c *Counter) Inc() {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Num++
}

func GoroutineInc() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	counter := new(Counter)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("counter.Num: %+v\n", counter.Num)
}

//题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ：原子操作、并发数据安全。

type LockFreeCounter struct {
	Num int64
}

func NewLockFreeCounter() *LockFreeCounter {
	return new(LockFreeCounter)
}

func (c *LockFreeCounter) Inc() {
	atomic.AddInt64(&c.Num, 1)
}

func (c *LockFreeCounter) Dec() {
	atomic.AddInt64(&c.Num, -1)
}

func GoroutineFreeLockCounter() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	lockFreeCounter := NewLockFreeCounter()

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				lockFreeCounter.Inc()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("lockFreeCounter.Num: %+v\n", lockFreeCounter.Num)
}
