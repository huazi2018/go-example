package main

import "fmt"
/**
这里我们创建了一个数组 a 来存放刚好 5 个 int。元素的类型和长度都是数组类型的一部分。数组默认是零值的，对于 int 数组来说也就是 0。

我们可以使用 array[index] = value 语法来设置数组指定位置的值，或者用 array[index] 得到值。

使用内置函数 len 返回数组的长度
使用这个语法在一行内初始化一个数组

数组的存储类型是单一的，但是你可以组合这些数据来构造多维的数据结构。
注意，在使用 fmt.Println 来打印数组的时候，会使用[v1 v2 v3 ...] 的格式显示
在典型的 Go 程序中，相对于数组而言，slice 使用的更多。我们将在后面讨论 slice。

 */
func main ()  {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
