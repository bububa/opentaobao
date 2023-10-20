package tmallnr

import (
	"sync"
)

// NewRetailResult 结构体
type NewRetailResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 请求结果
	ResultData *TagRespDto `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 成功或者失败
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}

var poolNewRetailResult = sync.Pool{
	New: func() any {
		return new(NewRetailResult)
	},
}

// GetNewRetailResult() 从对象池中获取NewRetailResult
func GetNewRetailResult() *NewRetailResult {
	return poolNewRetailResult.Get().(*NewRetailResult)
}

// ReleaseNewRetailResult 释放NewRetailResult
func ReleaseNewRetailResult(v *NewRetailResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.ResultData = nil
	v.SuccessFlag = false
	poolNewRetailResult.Put(v)
}
