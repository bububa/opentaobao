package yunosaccount

import (
	"sync"
)

// AccountResult 结构体
type AccountResult struct {
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// status
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolAccountResult = sync.Pool{
	New: func() any {
		return new(AccountResult)
	},
}

// GetAccountResult() 从对象池中获取AccountResult
func GetAccountResult() *AccountResult {
	return poolAccountResult.Get().(*AccountResult)
}

// ReleaseAccountResult 释放AccountResult
func ReleaseAccountResult(v *AccountResult) {
	v.Data = ""
	v.Message = ""
	v.Status = 0
	poolAccountResult.Put(v)
}
