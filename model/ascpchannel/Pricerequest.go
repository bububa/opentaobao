package ascpchannel

import (
	"sync"
)

// Pricerequest 结构体
type Pricerequest struct {
	// 产品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 价格类型
	PriceType string `json:"price_type,omitempty" xml:"price_type,omitempty"`
	// 经营模式
	SalesMode string `json:"sales_mode,omitempty" xml:"sales_mode,omitempty"`
	// 二级渠道code
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 市场code
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
}

var poolPricerequest = sync.Pool{
	New: func() any {
		return new(Pricerequest)
	},
}

// GetPricerequest() 从对象池中获取Pricerequest
func GetPricerequest() *Pricerequest {
	return poolPricerequest.Get().(*Pricerequest)
}

// ReleasePricerequest 释放Pricerequest
func ReleasePricerequest(v *Pricerequest) {
	v.ProductId = ""
	v.PriceType = ""
	v.SalesMode = ""
	v.SubChannelCode = ""
	v.ChannelCode = ""
	poolPricerequest.Put(v)
}
