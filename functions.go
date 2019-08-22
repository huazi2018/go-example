package main

import (
	"fmt"
)

/**
函数 是 Go 的中心。

这里是一个函数，接受两个 int 并且以 int 返回它们的和Go 需要明确的返回值，例如，它不会自动返回最后一个表达式的值
正如你期望的那样，通过 name(args) 来调用一个函数，
这里有许多 Go 函数的其他特性。其中一个就是多值返回，也是我们接下来需要接触的。
 */

func plus (a int, b int) int  {
	return a + b
}

func main ()  {
	res := plus(1, 2)
	fmt.Println("1+2=", res)
}
