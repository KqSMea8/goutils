package time

import (
	"time"

	"github.com/metakeule/fmtdate"
)

//包含end 天
func BuildTimeInterval(start, end time.Time) []string {
	startUnix := start.Unix()
	endUnix := end.Unix()
	length := (endUnix - startUnix) / 86400 //返回的是整数
	ret := make([]string, 0, length+1)      //只分配一次内存
	for i := startUnix; i < endUnix; i += 86400 {
		tmpday := fmtdate.Format(fmtdate.DefaultDateFormat, time.Unix(i, 0))
		ret = append(ret, tmpday)
	}
	return ret
}

//相减获取对应的time.Time
func GetBeforeDayTimer(day int64) time.Time {
	if day > 0 {
		day *= -1
	}
	return GetDayTimer(day)
}

//获取对应的天数
func GetDayTimer(day int64) time.Time {
	tunix := time.Now().Unix() + day*86400
	res := fmtdate.Format(fmtdate.DefaultDateFormat, time.Unix(tunix, 0))
	t, err := fmtdate.ParseDate(res)
	if err != nil {
		panic(err)
	}
	return t
}

//返回当天 00:00:00 对应的time.Time
func GetTodayTimer() time.Time {
	t, err := fmtdate.ParseDate(fmtdate.Format(fmtdate.DefaultDateFormat, time.Now()))
	if err != nil {
		panic(err)
	}
	return t
}

//获取昨天的Time
func GetYesterdayTimer() time.Time {
	return GetBeforeDayTimer(-1)
}
