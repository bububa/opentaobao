package mos

import (
	"sync"
)

// AlibabaMosOnsiteTradeRefundResultDo 结构体
type AlibabaMosOnsiteTradeRefundResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *RefundResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosOnsiteTradeRefundResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosOnsiteTradeRefundResultDo)
	},
}

// GetAlibabaMosOnsiteTradeRefundResultDo() 从对象池中获取AlibabaMosOnsiteTradeRefundResultDo
func GetAlibabaMosOnsiteTradeRefundResultDo() *AlibabaMosOnsiteTradeRefundResultDo {
	return poolAlibabaMosOnsiteTradeRefundResultDo.Get().(*AlibabaMosOnsiteTradeRefundResultDo)
}

// ReleaseAlibabaMosOnsiteTradeRefundResultDo 释放AlibabaMosOnsiteTradeRefundResultDo
func ReleaseAlibabaMosOnsiteTradeRefundResultDo(v *AlibabaMosOnsiteTradeRefundResultDo) {
	v.ErrMsg = ""
	v.Data = nil
	v.ErrCode = 0
	v.Success = false
	poolAlibabaMosOnsiteTradeRefundResultDo.Put(v)
}
