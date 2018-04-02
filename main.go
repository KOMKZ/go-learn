package main

import (
	"fmt"
	"time"
	"math/rand"
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
	
}