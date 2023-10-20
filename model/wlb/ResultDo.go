package wlb

import (
	"sync"
)

// ResultDo 结构体
type ResultDo struct {
	// 网络延时
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 01
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功、失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultDo = sync.Pool{
	New: func() any {
		return new(ResultDo)
	},
}

// GetResultDo() 从对象池中获取ResultDo
func GetResultDo() *ResultDo {
	return poolResultDo.Get().(*ResultDo)
}

// ReleaseResultDo 释放ResultDo
func ReleaseResultDo(v *ResultDo) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolResultDo.Put(v)
}
