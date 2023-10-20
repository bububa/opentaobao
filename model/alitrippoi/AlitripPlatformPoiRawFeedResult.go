package alitrippoi

import (
	"sync"
)

// AlitripPlatformPoiRawFeedResult 结构体
type AlitripPlatformPoiRawFeedResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回值
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// totalRecords
	TotalRecords int64 `json:"total_records,omitempty" xml:"total_records,omitempty"`
	// 成功标识
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPlatformPoiRawFeedResult = sync.Pool{
	New: func() any {
		return new(AlitripPlatformPoiRawFeedResult)
	},
}

// GetAlitripPlatformPoiRawFeedResult() 从对象池中获取AlitripPlatformPoiRawFeedResult
func GetAlitripPlatformPoiRawFeedResult() *AlitripPlatformPoiRawFeedResult {
	return poolAlitripPlatformPoiRawFeedResult.Get().(*AlitripPlatformPoiRawFeedResult)
}

// ReleaseAlitripPlatformPoiRawFeedResult 释放AlitripPlatformPoiRawFeedResult
func ReleaseAlitripPlatformPoiRawFeedResult(v *AlitripPlatformPoiRawFeedResult) {
	v.Message = ""
	v.ResultCode = ""
	v.Data = 0
	v.TotalRecords = 0
	v.Success = false
	poolAlitripPlatformPoiRawFeedResult.Put(v)
}
