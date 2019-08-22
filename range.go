package main

import "fmt"
/**
Range 遍历

range 迭代各种各样的数据结构。让我们来看看如何在我们已经学过的数据结构上使用 rang 吧。


 */
func main ()  {
	//这里我们使用 range 来统计一个 slice 的元素个数。数组也可以采用这种方法。
	nums := []int {2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//range 在数组和 slice 中都同样提供每个项的索引和值。上面我们不需要索引，所以我们使用 空值定义符_ 来忽略它。有时候我们实际上是需要这个索引的。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range 在 map 中迭代键值对。
	kvs := map[string]string {"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("%s -> %s\n", k, v)
	}

	//range 在字符串中迭代 unicode 编码。第一个返回值是rune 的起始字节位置，然后第二个是 rune 自己。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
