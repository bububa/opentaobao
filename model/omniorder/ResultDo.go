package omniorder

import (
	"sync"
)

// ResultDo 结构体
type ResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 订单详情
	Data *OrderDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
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
	v.ErrMsg = ""
	v.Data = nil
	v.ErrCode = 0
	v.Success = false
	poolResultDo.Put(v)
}
