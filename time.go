package main

import "fmt"
import "time"
/**
时间
 */
func main ()  {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2019,8, 25, 10, 45, 36, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p("============")
	//这些方法来比较两个时间，分别测试一下是否是之前，之后或者是同一时刻，精确到秒。
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	p("============")
	//方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间。
	diff := now.Sub(then)
	p(diff)

	//我们计算出不同单位下的时间长度值。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p("============")
	//你可以使用 Add 将时间后移一个时间间隔，或者使用一个 - 来将时间前移一个时间间隔。
	p(then.Add(diff))
	p(then.Add(-diff))
}
