package ascp

import (
	"sync"
)

// PurchasePrice 结构体
type PurchasePrice struct {
	// 币种，USD-美元，CNY-人民币，RUB-卢布，JPY-日元，EUR-欧元，GBP-英镑，HKD-港币，NZD-新西兰元，SGD-新加坡元，AUD-澳元，KRW-韩元，THB-泰铢
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 采购价，单位:分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolPurchasePrice = sync.Pool{
	New: func() any {
		return new(PurchasePrice)
	},
}

// GetPurchasePrice() 从对象池中获取PurchasePrice
func GetPurchasePrice() *PurchasePrice {
	return poolPurchasePrice.Get().(*PurchasePrice)
}

// ReleasePurchasePrice 释放PurchasePrice
func ReleasePurchasePrice(v *PurchasePrice) {
	v.Currency = ""
	v.SupplierName = ""
	v.Price = 0
	poolPurchasePrice.Put(v)
}
