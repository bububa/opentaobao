package ascpchannel

import (
	"sync"
)

// TopChannelPriceDetail 结构体
type TopChannelPriceDetail struct {
	// 币种
	CurrencyPriceValue string `json:"currency_price_value,omitempty" xml:"currency_price_value,omitempty"`
	// 价格值
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 扩展价格
	ExtendPrice string `json:"extend_price,omitempty" xml:"extend_price,omitempty"`
}

var poolTopChannelPriceDetail = sync.Pool{
	New: func() any {
		return new(TopChannelPriceDetail)
	},
}

// GetTopChannelPriceDetail() 从对象池中获取TopChannelPriceDetail
func GetTopChannelPriceDetail() *TopChannelPriceDetail {
	return poolTopChannelPriceDetail.Get().(*TopChannelPriceDetail)
}

// ReleaseTopChannelPriceDetail 释放TopChannelPriceDetail
func ReleaseTopChannelPriceDetail(v *TopChannelPriceDetail) {
	v.CurrencyPriceValue = ""
	v.Price = ""
	v.ExtendPrice = ""
	poolTopChannelPriceDetail.Put(v)
}
