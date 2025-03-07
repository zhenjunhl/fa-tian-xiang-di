package f

import (
	"time"
)

// GetTimestampHoursAgo 获取多少小时前的时间戳,十位
func GetTimestampHoursAgo(hours int) int64 {
	// 获取当前时间
	now := time.Now()
	// 计算指定小时数之前的时间
	pastTime := now.Add(-time.Duration(hours) * time.Hour)
	// 返回该时间的时间戳（秒级）
	return pastTime.Unix()
}
