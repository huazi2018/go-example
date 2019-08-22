package main

import "fmt"

/**
for 是 Go 中唯一的循环结构。这里有 for 循环的三个基本使用方式:

  最常用的方式，带单个循环条件;
  经典的初始化/条件/后续形式 for 循环;
  不带条件的 for 循环将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环。

  * 在教程后面，当我们学到 rang 语句，channels，以及其他数据结构时，将会看到一些 for 的其它使用形式
 */
func main ()  {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}
