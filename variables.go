package main

import (
	"fmt"
)

/**
 变量
 var 声明 1 个或者多个变量。
 你可以申明一次性声明多个变量。
 Go 将自动推断已经初始化的变量类型。
 声明变量且没有给出对应的初始值时，变量将会初始化为零值 。例如，一个 int 的零值是 0。
 := 语句是申明并初始化变量的简写，例如这个例子中的 var f string = "short"。
 */

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1,2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}