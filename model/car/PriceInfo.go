package car

import (
	"sync"
)

// PriceInfo 结构体
type PriceInfo struct {
	// 费用明细
	Detail []DetailPriceInfo `json:"detail,omitempty" xml:"detail>detail_price_info,omitempty"`
	// 总费用，折后金额总费用，折后金额  eg:203.00
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 原价，如果订单有折扣这里为折扣前的价格，如果没有折扣和totalPrice字段保持一致
	OriginalPrice string `json:"original_price,omitempty" xml:"original_price,omitempty"`
}

var poolPriceInfo = sync.Pool{
	New: func() any {
		return new(PriceInfo)
	},
}

// GetPriceInfo() 从对象池中获取PriceInfo
func GetPriceInfo() *PriceInfo {
	return poolPriceInfo.Get().(*PriceInfo)
}

// ReleasePriceInfo 释放PriceInfo
func ReleasePriceInfo(v *PriceInfo) {
	v.Detail = v.Detail[:0]
	v.TotalPrice = ""
	v.OriginalPrice = ""
	poolPriceInfo.Put(v)
}
