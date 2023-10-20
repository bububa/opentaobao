package xhotelcrm

import (
	"sync"
)

// MsdResult 结构体
type MsdResult struct {
	// 系统异常
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// S_SYS_EXCEPTION
	ErrCode bool `json:"err_code,omitempty" xml:"err_code,omitempty"`
}

var poolMsdResult = sync.Pool{
	New: func() any {
		return new(MsdResult)
	},
}

// GetMsdResult() 从对象池中获取MsdResult
func GetMsdResult() *MsdResult {
	return poolMsdResult.Get().(*MsdResult)
}

// ReleaseMsdResult 释放MsdResult
func ReleaseMsdResult(v *MsdResult) {
	v.ErrMsg = ""
	v.Success = false
	v.ErrCode = false
	poolMsdResult.Put(v)
}
