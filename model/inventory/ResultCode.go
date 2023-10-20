package inventory

import (
	"sync"
)

// ResultCode 结构体
type ResultCode struct {
	// 结果描述
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果码
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolResultCode = sync.Pool{
	New: func() any {
		return new(ResultCode)
	},
}

// GetResultCode() 从对象池中获取ResultCode
func GetResultCode() *ResultCode {
	return poolResultCode.Get().(*ResultCode)
}

// ReleaseResultCode 释放ResultCode
func ReleaseResultCode(v *ResultCode) {
	v.Code = ""
	v.Message = ""
	v.Id = 0
	poolResultCode.Put(v)
}
