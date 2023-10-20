package mtopopen

import (
	"sync"
)

// ActivityLotteryWriteResult 结构体
type ActivityLotteryWriteResult struct {
	// isv活动url
	H5Url string `json:"h5_url,omitempty" xml:"h5_url,omitempty"`
	// isv活动的id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

var poolActivityLotteryWriteResult = sync.Pool{
	New: func() any {
		return new(ActivityLotteryWriteResult)
	},
}

// GetActivityLotteryWriteResult() 从对象池中获取ActivityLotteryWriteResult
func GetActivityLotteryWriteResult() *ActivityLotteryWriteResult {
	return poolActivityLotteryWriteResult.Get().(*ActivityLotteryWriteResult)
}

// ReleaseActivityLotteryWriteResult 释放ActivityLotteryWriteResult
func ReleaseActivityLotteryWriteResult(v *ActivityLotteryWriteResult) {
	v.H5Url = ""
	v.ActivityId = 0
	poolActivityLotteryWriteResult.Put(v)
}
