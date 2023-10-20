package tblogistics

import (
	"sync"
)

// ResultWrappe 结构体
type ResultWrappe struct {
	// 错误编码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultWrappe = sync.Pool{
	New: func() any {
		return new(ResultWrappe)
	},
}

// GetResultWrappe() 从对象池中获取ResultWrappe
func GetResultWrappe() *ResultWrappe {
	return poolResultWrappe.Get().(*ResultWrappe)
}

// ReleaseResultWrappe 释放ResultWrappe
func ReleaseResultWrappe(v *ResultWrappe) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolResultWrappe.Put(v)
}
