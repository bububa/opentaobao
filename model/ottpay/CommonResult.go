package ottpay

import (
	"sync"
)

// CommonResult 结构体
type CommonResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 返回码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCommonResult = sync.Pool{
	New: func() any {
		return new(CommonResult)
	},
}

// GetCommonResult() 从对象池中获取CommonResult
func GetCommonResult() *CommonResult {
	return poolCommonResult.Get().(*CommonResult)
}

// ReleaseCommonResult 释放CommonResult
func ReleaseCommonResult(v *CommonResult) {
	v.Message = ""
	v.Data = nil
	v.Code = 0
	v.Success = false
	poolCommonResult.Put(v)
}
