package main

import (
	"fmt"
)

func main()  {
	// 无缓冲信道存消息和取消息都是阻塞的
	// 取消息会阻塞当前routine
	// 写如消息也会阻塞当前routine
	// 开启一个not buffered channel的时候 准备读取的时候 必须有数据进来 否则就会阻塞住一直等待数据写入
	// 对一个not buffered channel 进行写入操作的时候 如果数据没有被读出 则channel是满的 该routine一直会被挂起 知道数据被取出
	// 对于一个 not buffered channel 来说 读取大部分下是成对的

	// channel的作用可以用来通知main routine当前的routine已经完毕了 你可以退出了 例子如下
	/*
	ch := make(chan struct{})
	go func() {
		fmt.Println("handle 1 second...")
		time.Sleep(time.Second)
		ch <- struct{}{}
	}()
	<-ch
	fmt.Println("ok")
	// handle 1 second...
	// ok
	*/

	// 引起死锁 一个routine在等待数据 一个迟迟不肯写入数据 或者相反比如
	// 一般思维看下来 先写入到了ch 然后就把ch读出来非常完美不是吗？错
	// ch <- 1 写入的过程 如果ch的数据没有被取出去 这里就会被阻塞住了
	// 然后后面的的 <-ch 就一直在等着当前routine的释放 形成死锁deadlock
	/*
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
	*/

}