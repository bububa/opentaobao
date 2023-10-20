package user

import (
	"sync"
)

// UserResultCode 结构体
type UserResultCode struct {
	// 结果code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果类型
	ResultType string `json:"result_type,omitempty" xml:"result_type,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

var poolUserResultCode = sync.Pool{
	New: func() any {
		return new(UserResultCode)
	},
}

// GetUserResultCode() 从对象池中获取UserResultCode
func GetUserResultCode() *UserResultCode {
	return poolUserResultCode.Get().(*UserResultCode)
}

// ReleaseUserResultCode 释放UserResultCode
func ReleaseUserResultCode(v *UserResultCode) {
	v.ResultCode = ""
	v.ResultType = ""
	v.ResultMsg = ""
	poolUserResultCode.Put(v)
}
