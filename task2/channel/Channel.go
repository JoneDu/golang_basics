package channel

import (
	"fmt"
	"sync"
	"time"
)

//✅Channel
//题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
//考察点 ：通道的基本使用、协程间通信。

func ChannelNumTransfer() {
	wg := sync.WaitGroup{}
	// 初始化一个无缓冲channel
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 10 {
			ch <- i
		}
		close(ch)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Printf("channel.len:%v, value: %v \n", len(ch), i)
		}
	}()

	wg.Wait()
	fmt.Println("p finished")
}

func ChannelNumTraSelect() {
	ch := make(chan int)
	go func() {
		for i := range 10 {
			ch <- i
		}
	}()

	for {
		select {
		case i := <-ch:
			fmt.Println("receive", i)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout")
			return
		}
	}
}

//题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
//考察点 ：通道的缓冲机制。

func ChannelBuffer() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	count := 100
	ch := make(chan int, count)
	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			ch <- i + 1
		}
		close(ch)
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case i, closeFlag := <-ch:
				fmt.Println("receive:", i, ",len:", len(ch))
				if closeFlag && len(ch) == 0 {
					return
				}
			}
		}
	}()
	wg.Wait()
	fmt.Println("channelBuffer finished")
}
