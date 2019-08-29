package main

import (
	"fmt"
	"math/rand"
	"time"
)
/**
随机数
Go 的 math/rand 包提供了伪随机数生成器（英）。

 */
func main ()  {
	rand.Seed(142)
	answer := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	fmt.Println("Magic 8-Ball says:", answer[rand.Intn(len(answer))])

	//rand.Intn 返回一个随机的整数 n，0 <= n <= 100。
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	//rand.Float64 返回一个64位浮点数 f，0.0 <= f <= 1.0
	fmt.Println(rand.Float64())

	//这个技巧可以用来生成其他范围的随机浮点数，例如5.0 <= f <= 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print(rand.Float64() * 5 + 5)
	fmt.Println()

	//默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。要产生变化的序列，需要给定一个变化的种子。需要注意的是，如果你出于加密目的，需要使用随机数的话，请使用 crypto/rand 包，此方法不够安全。
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	//调用上面返回的 rand.Source 的函数和调用 rand 包中函数是相同的。
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	//如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))

}
