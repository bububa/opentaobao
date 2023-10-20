package aesolution

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// error message
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// time stamp
	TimeStamp string `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// error code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data *GlobalAeopTpOrderDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// success
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

var poolBaseResult = sync.Pool{
	New: func() any {
		return new(BaseResult)
	},
}

// GetBaseResult() 从对象池中获取BaseResult
func GetBaseResult() *BaseResult {
	return poolBaseResult.Get().(*BaseResult)
}

// ReleaseBaseResult 释放BaseResult
func ReleaseBaseResult(v *BaseResult) {
	v.ErrorMessage = ""
	v.TimeStamp = ""
	v.ErrorCode = ""
	v.Data = nil
	v.ResultSuccess = false
	poolBaseResult.Put(v)
}
