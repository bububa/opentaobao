package alsc

import (
	"sync"
)

// UnchargeOpenReq 结构体
type UnchargeOpenReq struct {
	// SaaS品牌ID(不能和outbrandid同时为空)
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 卡号
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 操作人ID(SaaS Id）
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 退储值订单Id，必填（kry是paymentItemId）,外部保证唯一，作为退款幂等号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// SaaS门店ID(不能和outshopid同时为空)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 原充值订单Id，必填（kry是srcPaymentItemId）
	SrcOuterOrderId string `json:"src_outer_order_id,omitempty" xml:"src_outer_order_id,omitempty"`
	// 外部门店ID(不能和sopid同时为空)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部品牌ID(不能和brandid同时为空)
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 请求id，幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// CS是辰森，KRY是客如云
	BizChannel string `json:"biz_channel,omitempty" xml:"biz_channel,omitempty"`
	// 原支付单号
	SrcOutPayId string `json:"src_out_pay_id,omitempty" xml:"src_out_pay_id,omitempty"`
}

var poolUnchargeOpenReq = sync.Pool{
	New: func() any {
		return new(UnchargeOpenReq)
	},
}

// GetUnchargeOpenReq() 从对象池中获取UnchargeOpenReq
func GetUnchargeOpenReq() *UnchargeOpenReq {
	return poolUnchargeOpenReq.Get().(*UnchargeOpenReq)
}

// ReleaseUnchargeOpenReq 释放UnchargeOpenReq
func ReleaseUnchargeOpenReq(v *UnchargeOpenReq) {
	v.BrandId = ""
	v.CardId = ""
	v.OperatorId = ""
	v.OuterOrderId = ""
	v.Remark = ""
	v.ShopId = ""
	v.SrcOuterOrderId = ""
	v.OutShopId = ""
	v.OutBrandId = ""
	v.RequestId = ""
	v.BizChannel = ""
	v.SrcOutPayId = ""
	poolUnchargeOpenReq.Put(v)
}
