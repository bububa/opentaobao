package fans

import (
	"sync"
)

// CashPoolVo 结构体
type CashPoolVo struct {
	// 付款url
	PayUrl string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
	// 奖金池id
	CashPoolId int64 `json:"cash_pool_id,omitempty" xml:"cash_pool_id,omitempty"`
}

var poolCashPoolVo = sync.Pool{
	New: func() any {
		return new(CashPoolVo)
	},
}

// GetCashPoolVo() 从对象池中获取CashPoolVo
func GetCashPoolVo() *CashPoolVo {
	return poolCashPoolVo.Get().(*CashPoolVo)
}

// ReleaseCashPoolVo 释放CashPoolVo
func ReleaseCashPoolVo(v *CashPoolVo) {
	v.PayUrl = ""
	v.CashPoolId = 0
	poolCashPoolVo.Put(v)
}
