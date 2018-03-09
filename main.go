package main

import (
	"fmt"
	"learn/ch4/treesort"
)

func main()  {
	s := []int{2,5,9,72,2,1,45}
	treesort.SortValues(s)
	fmt.Println(s)

}