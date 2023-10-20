package simba

import (
	"sync"
)

// TopAccountBalanceVo 结构体
type TopAccountBalanceVo struct {
	// 账户实时现金余额
	RtCashBalance string `json:"rt_cash_balance,omitempty" xml:"rt_cash_balance,omitempty"`
}

var poolTopAccountBalanceVo = sync.Pool{
	New: func() any {
		return new(TopAccountBalanceVo)
	},
}

// GetTopAccountBalanceVo() 从对象池中获取TopAccountBalanceVo
func GetTopAccountBalanceVo() *TopAccountBalanceVo {
	return poolTopAccountBalanceVo.Get().(*TopAccountBalanceVo)
}

// ReleaseTopAccountBalanceVo 释放TopAccountBalanceVo
func ReleaseTopAccountBalanceVo(v *TopAccountBalanceVo) {
	v.RtCashBalance = ""
	poolTopAccountBalanceVo.Put(v)
}
