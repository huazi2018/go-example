package main

import "fmt"
import "os"
/**
字符串格式化
 */
type point struct{
	x, y int
}

func main ()  {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)

	fmt.Printf("%#\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 33)

	fmt.Printf("%x\n", 456)
	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"strings\"")
	fmt.Printf("%q\n", "\"strings\"")

	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	//你可以使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout。
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
