package main

import (
	"learn/ch6"
	"fmt"
)

func main()  {





}
/**

	p := geometry.Point{1,1}
	q := geometry.Point{1,2}
	p.ScaleBy(3)
	q.ScaleBy(3)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", q)

	p := geometry.Point{2,3}
	q := geometry.Point{3,5}
	// 下面这句compile出错 因为没有地址
	(geometry.Point{2,3}).ScaleBy(4)
	fmt.Println(p.Distance(q))

	// 指针类型方法的使用
	p := geometry.Point{2,3}
	pp := &p
	q := geometry.Point{3,4}
	fmt.Println(pp.Distance(q))
	(pp).ScaleBy(4)
	fmt.Println(pp.Distance(q))

	// 指针类型方法的使用
	p := &geometry.Point{2,3}
	q := geometry.Point{3,4}
	fmt.Println(p.Distance(q))
	(p).ScaleBy(4)
	fmt.Println(p.Distance(q))

	// 指针类型方法的使用
	p := geometry.Point{2,3}
	q := geometry.Point{3,4}
	fmt.Println(p.Distance(q))
	(&p).ScaleBy(4)
	fmt.Println(p.Distance(q))

	// package 方法以及 类型的method能够同名
	p1 := geometry.Path{
		{2,3},
		{2,3},
		{3,3},
	}
	fmt.Printf("method sum is: %f\n", p1.Distance())
	fmt.Printf("package sum is: %f\n", geometry.PathDistance(p1))

	// 定义一个Path计算其中所有点的距离
	p1 := geometry.Path{
		{2,3},
		{2,3},
		{3,3},
	}
	fmt.Printf("sum is: %f\n", p1.Distance())

	// 计算两个点的距离
	p := geometry.Point{1,2}
	q := geometry.Point{2,3}
	fmt.Println(p.Distance(q))
 */
