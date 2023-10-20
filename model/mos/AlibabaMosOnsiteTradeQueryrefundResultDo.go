package mos

import (
	"sync"
)

// AlibabaMosOnsiteTradeQueryrefundResultDo 结构体
type AlibabaMosOnsiteTradeQueryrefundResultDo struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *OnsiteRefundResponse `json:"data,omitempty" xml:"data,omitempty"`
	// errCode
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosOnsiteTradeQueryrefundResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosOnsiteTradeQueryrefundResultDo)
	},
}

// GetAlibabaMosOnsiteTradeQueryrefundResultDo() 从对象池中获取AlibabaMosOnsiteTradeQueryrefundResultDo
func GetAlibabaMosOnsiteTradeQueryrefundResultDo() *AlibabaMosOnsiteTradeQueryrefundResultDo {
	return poolAlibabaMosOnsiteTradeQueryrefundResultDo.Get().(*AlibabaMosOnsiteTradeQueryrefundResultDo)
}

// ReleaseAlibabaMosOnsiteTradeQueryrefundResultDo 释放AlibabaMosOnsiteTradeQueryrefundResultDo
func ReleaseAlibabaMosOnsiteTradeQueryrefundResultDo(v *AlibabaMosOnsiteTradeQueryrefundResultDo) {
	v.ErrMsg = ""
	v.Data = nil
	v.ErrCode = 0
	v.Success = false
	poolAlibabaMosOnsiteTradeQueryrefundResultDo.Put(v)
}
