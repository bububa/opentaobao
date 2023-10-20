package tvpay

import (
	"sync"
)

// QueryPayResultResultDo 结构体
type QueryPayResultResultDo struct {
	// 支付资金组成情况
	FundMoney string `json:"fund_money,omitempty" xml:"fund_money,omitempty"`
	// 支付资金组成情况
	FundMoneyCode string `json:"fund_money_code,omitempty" xml:"fund_money_code,omitempty"`
	// 订单状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolQueryPayResultResultDo = sync.Pool{
	New: func() any {
		return new(QueryPayResultResultDo)
	},
}

// GetQueryPayResultResultDo() 从对象池中获取QueryPayResultResultDo
func GetQueryPayResultResultDo() *QueryPayResultResultDo {
	return poolQueryPayResultResultDo.Get().(*QueryPayResultResultDo)
}

// ReleaseQueryPayResultResultDo 释放QueryPayResultResultDo
func ReleaseQueryPayResultResultDo(v *QueryPayResultResultDo) {
	v.FundMoney = ""
	v.FundMoneyCode = ""
	v.Status = ""
	poolQueryPayResultResultDo.Put(v)
}
