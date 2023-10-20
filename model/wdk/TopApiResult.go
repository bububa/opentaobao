package wdk

import (
	"sync"
)

// TopApiResult 结构体
type TopApiResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误详情
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 数据
	Data *WarehouseNetworkResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 请求成功或失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopApiResult = sync.Pool{
	New: func() any {
		return new(TopApiResult)
	},
}

// GetTopApiResult() 从对象池中获取TopApiResult
func GetTopApiResult() *TopApiResult {
	return poolTopApiResult.Get().(*TopApiResult)
}

// ReleaseTopApiResult 释放TopApiResult
func ReleaseTopApiResult(v *TopApiResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolTopApiResult.Put(v)
}
