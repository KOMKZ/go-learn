package main

import (
	"time"
	"fmt"
	"strings"
)

func main()  {
	// loading 放在另一个 gorouting 中来执行
	go spinner(1000 * time.Microsecond)
	const n = 45
	// 主程序退出的时候，所有的gorouting跟着终端
	fmt.Printf("\rFibonacci(%d) = %d", n, fib(n)) // slow
}

func  fib(x int) int  {
	if x < 2{
		return x
	}
	return fib(x -1 ) + fib(x -2 )
}

func spinner(delay time.Duration)  {
	for {
		for i, s := range "...." {
			// \r 能够清除这些当前的输入
			fmt.Printf("\rloading%s", strings.Repeat(string(s), i + 1))
			time.Sleep(delay)
		}
	}
}

