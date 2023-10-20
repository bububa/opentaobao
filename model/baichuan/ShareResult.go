package baichuan

import (
	"sync"
)

// ShareResult 结构体
type ShareResult struct {
	// model
	Model *PasswordRuleResultDto `json:"model,omitempty" xml:"model,omitempty"`
	// resultCode
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// totalNumber
	TotalNumber int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolShareResult = sync.Pool{
	New: func() any {
		return new(ShareResult)
	},
}

// GetShareResult() 从对象池中获取ShareResult
func GetShareResult() *ShareResult {
	return poolShareResult.Get().(*ShareResult)
}

// ReleaseShareResult 释放ShareResult
func ReleaseShareResult(v *ShareResult) {
	v.Model = nil
	v.ResultCode = nil
	v.TotalNumber = 0
	v.IsSuccess = false
	poolShareResult.Put(v)
}
