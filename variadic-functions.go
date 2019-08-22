package main

import "fmt"

/**
可变参数函数。可以用任意数量的参数调用。例如，fmt.Println 是一个常见的变参函数。

Go 函数的另一个关键的方面是闭包结构，这是接下来我们需要看看的。
 */

//这个函数使用任意数目的 int 作为参数。
func sum (nums ...int)  {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main ()  {
	sum(1, 2)
	sum(1, 2, 3)

	//如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)。
	nums := []int{1, 2, 3}
	sum(nums...)
}