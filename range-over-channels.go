package main

import "fmt"
/**
通道遍历
在前面的例子中，我们讲过 for 和 range为基本的数据结构提供了迭代的功能。我们也可以使用这个语法来遍历从通道中取得的值。
这个例子也让我们看到，一个非空的通道也是可以关闭的，但是通道中剩下的值仍然可以被接收到
 */
func main ()  {
	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"
	close(queue)

	//这个 range 迭代从 queue 中得到的每个值。因为我们在前面 close 了这个通道，这个迭代会在接收完 2 个值之后结束。
	// 如果我们没有 close 它，我们将在这个循环中继续阻塞执行，等待接收第三个值
	for elem := range queue {
		fmt.Println(elem)
	}
}