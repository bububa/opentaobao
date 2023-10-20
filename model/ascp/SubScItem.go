package ascp

import (
	"sync"
)

// SubScItem 结构体
type SubScItem struct {
	// 子货品ERP货品id
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 子货品商家编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 币种，USD-美元，CNY-人民币，RUB-卢布，JPY-日元，EUR-欧元，GBP-英镑，HKD-港币，NZD-新西兰元，SGD-新加坡元，AUD-澳元，KRW-韩元，THB-泰铢
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 子货品数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 子货品零售价（人民币-分）
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 固定价格，1=是，0=否
	FixedPrice int64 `json:"fixed_price,omitempty" xml:"fixed_price,omitempty"`
}

var poolSubScItem = sync.Pool{
	New: func() any {
		return new(SubScItem)
	},
}

// GetSubScItem() 从对象池中获取SubScItem
func GetSubScItem() *SubScItem {
	return poolSubScItem.Get().(*SubScItem)
}

// ReleaseSubScItem 释放SubScItem
func ReleaseSubScItem(v *SubScItem) {
	v.ScItemId = ""
	v.ScItemCode = ""
	v.Currency = ""
	v.Quantity = 0
	v.RetailPrice = 0
	v.FixedPrice = 0
	poolSubScItem.Put(v)
}
