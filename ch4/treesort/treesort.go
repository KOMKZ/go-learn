package treesort

import "fmt"

// 数结构用于存储排序的数据
type tree struct {
	value int
	left, right *tree
}

// 增加一个值到树中
func add(t *tree, v int) *tree{
	if nil == t {
		t = new(tree)
		t.value = v
		return t
	}
	if t.value > v {
		t.left = add(t.left, v)
	}else{
		t.right = add(t.right, v)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if nil != t {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	fmt.Println(values)
	return values
}

func SortValues(values []int){
	var tree *tree
	for _, v := range values{
		tree = add(tree, v)
	}
	appendValues(values[:0], tree)
}

