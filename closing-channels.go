package main

import "fmt"
/**
通道的关闭
关闭 一个通道意味着不能再向这个通道发送值了。这个特性可以用来给这个通道的接收方传达工作已经完成的信息。

在这个例子中，我们将使用一个 jobs 通道来传递 main() 中 Go协程任务执行的结束信息到一个工作 Go 协程中。
当我们没有多余的任务给这个工作 Go 协程时，我们将 close 这个 jobs 通道。


 */
func main ()  {
	jobs := make(chan int, 5)
	done := make(chan bool)

	/*
	这是工作 Go 协程。使用 j, more := <- jobs 循环的从jobs 接收数据。在接收的这个特殊的二值形式的值中，如果 jobs 已经关闭了，
	并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。当我们完成所有的任务时，将使用这个特性通过 done 通道去进行通知。
	 */
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	//这里使用 jobs 发送 3 个任务到工作函数中，然后关闭 jobs。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job")
	}
	close(jobs)
	fmt.Println("send all jobs")

	<- done
}
