package icbulogistics

import (
	"sync"
)

// ExpressQuoteItemList 结构体
type ExpressQuoteItemList struct {
	// 费用编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 费用名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 费用描述
	ChargeDesc string `json:"charge_desc,omitempty" xml:"charge_desc,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 费用类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 价格信息
	SalesAmount *Money `json:"sales_amount,omitempty" xml:"sales_amount,omitempty"`
}

var poolExpressQuoteItemList = sync.Pool{
	New: func() any {
		return new(ExpressQuoteItemList)
	},
}

// GetExpressQuoteItemList() 从对象池中获取ExpressQuoteItemList
func GetExpressQuoteItemList() *ExpressQuoteItemList {
	return poolExpressQuoteItemList.Get().(*ExpressQuoteItemList)
}

// ReleaseExpressQuoteItemList 释放ExpressQuoteItemList
func ReleaseExpressQuoteItemList(v *ExpressQuoteItemList) {
	v.Code = ""
	v.Name = ""
	v.ChargeDesc = ""
	v.Currency = ""
	v.Type = ""
	v.Quantity = 0
	v.SalesAmount = nil
	poolExpressQuoteItemList.Put(v)
}
