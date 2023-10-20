package alsc

import (
	"sync"
)

// ExtraConsumePointOpenReq 结构体
type ExtraConsumePointOpenReq struct {
	// CS是辰森，KRY是客如云
	BizChannel string `json:"biz_channel,omitempty" xml:"biz_channel,omitempty"`
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 顾客id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 操作人ID
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 关联交易号/订单号
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 变更理由
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 请求id，用来做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部门店ID,shop_id和out_shop_id不可同时为空
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部品牌id,brandId与out_brand_id不可同时为空
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 变更积分数
	ChangePoint int64 `json:"change_point,omitempty" xml:"change_point,omitempty"`
}

var poolExtraConsumePointOpenReq = sync.Pool{
	New: func() any {
		return new(ExtraConsumePointOpenReq)
	},
}

// GetExtraConsumePointOpenReq() 从对象池中获取ExtraConsumePointOpenReq
func GetExtraConsumePointOpenReq() *ExtraConsumePointOpenReq {
	return poolExtraConsumePointOpenReq.Get().(*ExtraConsumePointOpenReq)
}

// ReleaseExtraConsumePointOpenReq 释放ExtraConsumePointOpenReq
func ReleaseExtraConsumePointOpenReq(v *ExtraConsumePointOpenReq) {
	v.BizChannel = ""
	v.BrandId = ""
	v.CustomerId = ""
	v.OperatorId = ""
	v.OutBizId = ""
	v.Reason = ""
	v.RequestId = ""
	v.ShopId = ""
	v.OutShopId = ""
	v.OutBrandId = ""
	v.ChangePoint = 0
	poolExtraConsumePointOpenReq.Put(v)
}
