package wdk

import (
	"sync"
)

// WorkResult 结构体
type WorkResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 返回数据
	Data *PackageResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 响应结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 返回结果
	ResultData bool `json:"result_data,omitempty" xml:"result_data,omitempty"`
}

var poolWorkResult = sync.Pool{
	New: func() any {
		return new(WorkResult)
	},
}

// GetWorkResult() 从对象池中获取WorkResult
func GetWorkResult() *WorkResult {
	return poolWorkResult.Get().(*WorkResult)
}

// ReleaseWorkResult 释放WorkResult
func ReleaseWorkResult(v *WorkResult) {
	v.Message = ""
	v.Code = ""
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Data = nil
	v.Success = false
	v.IsSuccess = false
	v.ResultData = false
	poolWorkResult.Put(v)
}
