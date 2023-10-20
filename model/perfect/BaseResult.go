package perfect

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// 返回的执行状态吗
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的数据实体
	Data *PerfectPerformanceItemPublishResp `json:"data,omitempty" xml:"data,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
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
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolBaseResult.Put(v)
}
