package main

import "fmt"

/**
 值
Go 拥有各值类型，包括字符串，整形，浮点型，布尔型等。下面是一些基本的例子。
 */
 
func main() {
    //字符串可以通过 + 连接。
    fmt.Println("go" + "lang")

	//整数和浮点数
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

	//布尔型，还有你想要的逻辑运算符。
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}