package interact

import (
	"sync"
)

// ActivityWriteResult 结构体
type ActivityWriteResult struct {
	// 活动页面h5链接
	H5Url string `json:"h5_url,omitempty" xml:"h5_url,omitempty"`
	// 报名成功的id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

var poolActivityWriteResult = sync.Pool{
	New: func() any {
		return new(ActivityWriteResult)
	},
}

// GetActivityWriteResult() 从对象池中获取ActivityWriteResult
func GetActivityWriteResult() *ActivityWriteResult {
	return poolActivityWriteResult.Get().(*ActivityWriteResult)
}

// ReleaseActivityWriteResult 释放ActivityWriteResult
func ReleaseActivityWriteResult(v *ActivityWriteResult) {
	v.H5Url = ""
	v.ActivityId = 0
	poolActivityWriteResult.Put(v)
}
