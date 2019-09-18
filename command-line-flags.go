package main

import "flag"
import "fmt"

/*
 * 命令行标志
 * 命令行标志是命令行程序指定选项的常用方式。例如，在 wc -l 中，这个 -l 就是一个命令行标志
 */
func main() {
	wordPtr := flag.String("word", "foo", "a string")
	
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	
	flag.Parse()
	
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}