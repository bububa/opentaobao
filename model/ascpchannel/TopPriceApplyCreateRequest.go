package ascpchannel

import (
	"sync"
)

// TopPriceApplyCreateRequest 结构体
type TopPriceApplyCreateRequest struct {
	// 操作人名称
	CreatorNick string `json:"creator_nick,omitempty" xml:"creator_nick,omitempty"`
	// 最低限价
	AdviseSalePriceLow string `json:"advise_sale_price_low,omitempty" xml:"advise_sale_price_low,omitempty"`
	// 操作人id
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// SKU渠道价
	SkuPriceMap string `json:"sku_price_map,omitempty" xml:"sku_price_map,omitempty"`
	// 建议零售价
	AdviseSalePriceHigh string `json:"advise_sale_price_high,omitempty" xml:"advise_sale_price_high,omitempty"`
	// 分销商名称，设置专享价填写
	DistributorNick string `json:"distributor_nick,omitempty" xml:"distributor_nick,omitempty"`
	// 渠道价格(含税)
	DistributePrice string `json:"distribute_price,omitempty" xml:"distribute_price,omitempty"`
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 渠道价格(未税)
	ProductPrice string `json:"product_price,omitempty" xml:"product_price,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 渠道产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolTopPriceApplyCreateRequest = sync.Pool{
	New: func() any {
		return new(TopPriceApplyCreateRequest)
	},
}

// GetTopPriceApplyCreateRequest() 从对象池中获取TopPriceApplyCreateRequest
func GetTopPriceApplyCreateRequest() *TopPriceApplyCreateRequest {
	return poolTopPriceApplyCreateRequest.Get().(*TopPriceApplyCreateRequest)
}

// ReleaseTopPriceApplyCreateRequest 释放TopPriceApplyCreateRequest
func ReleaseTopPriceApplyCreateRequest(v *TopPriceApplyCreateRequest) {
	v.CreatorNick = ""
	v.AdviseSalePriceLow = ""
	v.CreatorId = ""
	v.SkuPriceMap = ""
	v.AdviseSalePriceHigh = ""
	v.DistributorNick = ""
	v.DistributePrice = ""
	v.SubChannelCode = ""
	v.ProductPrice = ""
	v.ChannelCode = ""
	v.ProductId = 0
	poolTopPriceApplyCreateRequest.Put(v)
}
