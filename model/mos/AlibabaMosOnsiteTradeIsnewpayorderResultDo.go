package mos

import (
	"sync"
)

// AlibabaMosOnsiteTradeIsnewpayorderResultDo 结构体
type AlibabaMosOnsiteTradeIsnewpayorderResultDo struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// errCode
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否为新支付订单。true：是，false:不是
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosOnsiteTradeIsnewpayorderResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosOnsiteTradeIsnewpayorderResultDo)
	},
}

// GetAlibabaMosOnsiteTradeIsnewpayorderResultDo() 从对象池中获取AlibabaMosOnsiteTradeIsnewpayorderResultDo
func GetAlibabaMosOnsiteTradeIsnewpayorderResultDo() *AlibabaMosOnsiteTradeIsnewpayorderResultDo {
	return poolAlibabaMosOnsiteTradeIsnewpayorderResultDo.Get().(*AlibabaMosOnsiteTradeIsnewpayorderResultDo)
}

// ReleaseAlibabaMosOnsiteTradeIsnewpayorderResultDo 释放AlibabaMosOnsiteTradeIsnewpayorderResultDo
func ReleaseAlibabaMosOnsiteTradeIsnewpayorderResultDo(v *AlibabaMosOnsiteTradeIsnewpayorderResultDo) {
	v.ErrMsg = ""
	v.ErrCode = 0
	v.Data = false
	v.Success = false
	poolAlibabaMosOnsiteTradeIsnewpayorderResultDo.Put(v)
}
