package maitix

import (
	"sync"
)

// QueryPayResultParam 结构体
type QueryPayResultParam struct {
	// 订单金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 订单日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 大麦订单号
	DamaiOrderNo int64 `json:"damai_order_no,omitempty" xml:"damai_order_no,omitempty"`
}

var poolQueryPayResultParam = sync.Pool{
	New: func() any {
		return new(QueryPayResultParam)
	},
}

// GetQueryPayResultParam() 从对象池中获取QueryPayResultParam
func GetQueryPayResultParam() *QueryPayResultParam {
	return poolQueryPayResultParam.Get().(*QueryPayResultParam)
}

// ReleaseQueryPayResultParam 释放QueryPayResultParam
func ReleaseQueryPayResultParam(v *QueryPayResultParam) {
	v.Amount = ""
	v.Date = ""
	v.OrderNo = ""
	v.DamaiOrderNo = 0
	poolQueryPayResultParam.Put(v)
}
