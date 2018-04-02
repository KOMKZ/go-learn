package main

import (
	"fmt"
)

/*
func xrange() chan int{
	ch := make(chan int)
	go func() {
		for i:=0;;i++{
			ch <- i
		}
	}()
	return ch
}

func get_msg(user string) chan string{
	ch := make(chan string)
	go func() {
		ch <- fmt.Sprintf("%s, Hello World\n", user)
	}()
	return ch
}
*/
/*
func do_stuff(i int)int  {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	return 100 - i
}
func branch(i int) chan int{
	ch := make(chan int)
	go func() {
		ch <- do_stuff(i)
	}()
	return ch
}
func fanIn(chs... chan int) chan int{
	ch := make(chan int)
	for _,c := range chs{
		go func(c chan int) {
			//
			ch <- <- c
		}(c)
	}
	return ch
}
*/

func foo(i int)chan int{
	ch := make(chan int)
	go func() {ch <- i}()
	return ch
}

func main()  {
	// 应用-生成器 模拟python的xrange 只有登到数据需要的时候才生成数据adf
	/*
	generator := xrange()
	for i:=0; i < 1000; i++{
		fmt.Println(<-generator)
	}
	*/
	// 应用-服务话
	/*
	jack := get_msg("jack")
	lartik := get_msg("lartik")
	fmt.Println(<-jack)
	fmt.Println(<-lartik)
	*/
	// 应用 多路复合 可以将多个channel的数据合并到一个数据
	/*
	result := fanIn(branch(1), branch(2), branch(3))
	for i:=0; i<3; i++ {
		fmt.Println(<-result)
	}
	*/

	// 应用 select 监听信道 同时收集多个管道的数据
	ch1, ch2, ch3 := foo(1), foo(2), foo(3)
	ch := make(chan int)
	go func() {
		for  {
			select {
			case v1:=<-ch1:ch<-v1
			case v2:=<-ch2:ch<-v2
			case v3:=<-ch3:ch<-v3
			}
		}
	}()
	for i:=0; i<3;i++{
		fmt.Println(<-ch)
	}
	fmt.Println("over")
}