package logistic

import (
	"sync"
)

// LogisticsResult 结构体
type LogisticsResult struct {
	// 错误编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 数据
	Data *Pagination `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolLogisticsResult = sync.Pool{
	New: func() any {
		return new(LogisticsResult)
	},
}

// GetLogisticsResult() 从对象池中获取LogisticsResult
func GetLogisticsResult() *LogisticsResult {
	return poolLogisticsResult.Get().(*LogisticsResult)
}

// ReleaseLogisticsResult 释放LogisticsResult
func ReleaseLogisticsResult(v *LogisticsResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolLogisticsResult.Put(v)
}
