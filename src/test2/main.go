package main

import "fmt"

// Sortable : 插入排序的抽象方法集
type Sortable interface {
	Len() int
	Less(i int, j int) bool
	Swap(i int, j int)
}

// InsertSort : 插入排序算法
func InsertSort(vec Sortable) {
	var l = vec.Len()

	var i = 1
	for {
		if i >= l {
			break
		}
		var j = i - 1
		for {
			if j < 0 {
				break
			}

			if vec.Less(j, j+1) {
				break
			}

			vec.Swap(j, j+1)

			j--
		}

		i++
	}
}

// Int32Sortable : int32 的排序的接口
type Int32Sortable []int32

// Len ： 序列长度
func (seq Int32Sortable) Len() int {
	return len(seq)
}

// Less : 比较算法
func (seq Int32Sortable) Less(i int, j int) bool {
	if seq[i] < seq[j] {
		return true
	}

	return false
}

// Swap : 交换算法
func (seq Int32Sortable) Swap(i int, j int) {
	seq[i], seq[j] = seq[j], seq[i]
}

func main() {
	var nums = []int32{2, 3, 1, 4, 9, 6, 7, 8, 5}

	fmt.Println(nums)

	InsertSort(Int32Sortable(nums))

	fmt.Println(nums)
}
