package alsc

import (
	"sync"
)

// QueryRechargeAccountOpenReq 结构体
type QueryRechargeAccountOpenReq struct {
	// 账户Id，如果有，则查询单个账户
	AccountId string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 品牌ID(不能和outbrandid同时为空)
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 卡Id，POS下查询会员卡/礼品卡在门店下储值账户
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 会员Id,查询会员下所有储值账户
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 门店ID(不能和outshopid同时为空))
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部门店ID(不能和shopid同时为空)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部品牌ID(不能和brandid同时为空)
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
}

var poolQueryRechargeAccountOpenReq = sync.Pool{
	New: func() any {
		return new(QueryRechargeAccountOpenReq)
	},
}

// GetQueryRechargeAccountOpenReq() 从对象池中获取QueryRechargeAccountOpenReq
func GetQueryRechargeAccountOpenReq() *QueryRechargeAccountOpenReq {
	return poolQueryRechargeAccountOpenReq.Get().(*QueryRechargeAccountOpenReq)
}

// ReleaseQueryRechargeAccountOpenReq 释放QueryRechargeAccountOpenReq
func ReleaseQueryRechargeAccountOpenReq(v *QueryRechargeAccountOpenReq) {
	v.AccountId = ""
	v.BrandId = ""
	v.CardId = ""
	v.CustomerId = ""
	v.ShopId = ""
	v.OutShopId = ""
	v.OutBrandId = ""
	poolQueryRechargeAccountOpenReq.Put(v)
}
