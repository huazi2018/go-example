package main

import (
	"fmt"
	"time"
)
/**
时间戳
一般程序会有获取 Unix 时间的秒数，毫秒数，或者微秒数的需要。

 */
func main ()  {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	//你也可以将协调世界时起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

