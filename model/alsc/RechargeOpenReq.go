package alsc

import (
	"sync"
)

// RechargeOpenReq 结构体
type RechargeOpenReq struct {
	// 时间，选填，不填就以平台为准
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 卡Id，礼品卡或会员卡Id
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 充值类型，必填
	ChargeType string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// SaaS品牌ID(不能和outbrandid同时为空)
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 操作人ID(SaaS Id)
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// SaaS门店ID(不能和outshopid同时为空)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部交易单号id(外部调用方保证在isv内部是唯一，可以是paymentItemId)
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 幂等号
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 外部支付单id
	OutPayId string `json:"out_pay_id,omitempty" xml:"out_pay_id,omitempty"`
	// 外部门店id(不能和shopid同时为空)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部品牌ID(不能和brandid同时为空)
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// CS是辰森，KRY是客如云
	BizChannel string `json:"biz_channel,omitempty" xml:"biz_channel,omitempty"`
	// 支付方式需要按照标准格式传入
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 赠储值
	GiftValue int64 `json:"gift_value,omitempty" xml:"gift_value,omitempty"`
	// 预储
	PreValue int64 `json:"pre_value,omitempty" xml:"pre_value,omitempty"`
	// 实储值
	RealValue int64 `json:"real_value,omitempty" xml:"real_value,omitempty"`
	// 是否触发储值赠送，必填    1. true：触发赠送，realValue必填      2. false：不触发赠送，realValue和giftValue必填
	TriggerGift bool `json:"trigger_gift,omitempty" xml:"trigger_gift,omitempty"`
}

var poolRechargeOpenReq = sync.Pool{
	New: func() any {
		return new(RechargeOpenReq)
	},
}

// GetRechargeOpenReq() 从对象池中获取RechargeOpenReq
func GetRechargeOpenReq() *RechargeOpenReq {
	return poolRechargeOpenReq.Get().(*RechargeOpenReq)
}

// ReleaseRechargeOpenReq 释放RechargeOpenReq
func ReleaseRechargeOpenReq(v *RechargeOpenReq) {
	v.BizDate = ""
	v.CardId = ""
	v.ChargeType = ""
	v.BrandId = ""
	v.OperatorId = ""
	v.ShopId = ""
	v.OuterOrderId = ""
	v.Remark = ""
	v.RequestId = ""
	v.OutPayId = ""
	v.OutShopId = ""
	v.OutBrandId = ""
	v.BizChannel = ""
	v.ExtInfo = ""
	v.GiftValue = 0
	v.PreValue = 0
	v.RealValue = 0
	v.TriggerGift = false
	poolRechargeOpenReq.Put(v)
}
