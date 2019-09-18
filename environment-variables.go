package main

import (
   "os"
   "strings"
   "fmt"
)
/*
 环境变量
 环境变量是一个在为 Unix 程序传递配置信息的普遍方式。让我们来看看如何设置，获取并列举环境变量。
*/

func main() {
	//使用 os.Setenv 来设置一个键值队
	os.Setenv("FOO", "1")
	
	//使用 os.Getenv获取一个键对应的值。如果键不存在，将会返回一个空字符串。
	fmt.Println("Foo:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	
	fmt.Println()
	
	//使用 os.Environ 来列出所有环境变量键值队。
	//这个函数会返回一个 KEY=value 形式的字符串切片。
	//你可以使用strings.Split 来得到键和值。这里我们打印所有的键
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}
}