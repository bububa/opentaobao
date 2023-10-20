package product

import (
	"sync"
)

// ResultDo 结构体
type ResultDo struct {
	// 错误码集合
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// 实际返回值
	Model *ErrorCode `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
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
	v.ErrorCodes = v.ErrorCodes[:0]
	v.Model = nil
	v.Success = false
	poolResultDo.Put(v)
}
