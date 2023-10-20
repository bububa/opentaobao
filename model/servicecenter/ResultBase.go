package servicecenter

import (
	"sync"
)

// ResultBase 结构体
type ResultBase struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// value
	Value *TpFundsRecoverResultDo `json:"value,omitempty" xml:"value,omitempty"`
	// errorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 查询接口是否OK
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultBase = sync.Pool{
	New: func() any {
		return new(ResultBase)
	},
}

// GetResultBase() 从对象池中获取ResultBase
func GetResultBase() *ResultBase {
	return poolResultBase.Get().(*ResultBase)
}

// ReleaseResultBase 释放ResultBase
func ReleaseResultBase(v *ResultBase) {
	v.ErrorMsg = ""
	v.Value = nil
	v.ErrorCode = 0
	v.Success = false
	poolResultBase.Put(v)
}
