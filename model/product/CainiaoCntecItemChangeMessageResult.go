package product

import (
	"sync"
)

// CainiaoCntecItemChangeMessageResult 结构体
type CainiaoCntecItemChangeMessageResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否成功接受到请求
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

var poolCainiaoCntecItemChangeMessageResult = sync.Pool{
	New: func() any {
		return new(CainiaoCntecItemChangeMessageResult)
	},
}

// GetCainiaoCntecItemChangeMessageResult() 从对象池中获取CainiaoCntecItemChangeMessageResult
func GetCainiaoCntecItemChangeMessageResult() *CainiaoCntecItemChangeMessageResult {
	return poolCainiaoCntecItemChangeMessageResult.Get().(*CainiaoCntecItemChangeMessageResult)
}

// ReleaseCainiaoCntecItemChangeMessageResult 释放CainiaoCntecItemChangeMessageResult
func ReleaseCainiaoCntecItemChangeMessageResult(v *CainiaoCntecItemChangeMessageResult) {
	v.ErrCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Model = false
	poolCainiaoCntecItemChangeMessageResult.Put(v)
}
