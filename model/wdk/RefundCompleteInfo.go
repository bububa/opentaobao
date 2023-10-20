package wdk

import (
	"sync"
)

// RefundCompleteInfo 结构体
type RefundCompleteInfo struct {
	// 子单
	SubRefundOrders []SubRefundOrder `json:"sub_refund_orders,omitempty" xml:"sub_refund_orders>sub_refund_order,omitempty"`
	// 支付渠道
	PayChannels []PayChannel `json:"pay_channels,omitempty" xml:"pay_channels>pay_channel,omitempty"`
	// 外部主单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 外部渠道店ID(与shop_id必选其一)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部逆向单ID
	OutRefundId string `json:"out_refund_id,omitempty" xml:"out_refund_id,omitempty"`
	// 渠道店id(与out_shop_id必选其一)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 退的商品费
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退的运费
	RefundPostFee int64 `json:"refund_post_fee,omitempty" xml:"refund_post_fee,omitempty"`
	// 退的包装费
	RefundPackageFee int64 `json:"refund_package_fee,omitempty" xml:"refund_package_fee,omitempty"`
	// 订单来源(选填out_shop_id时该值必填)
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 渠道退商家的佣金，单位：分
	Commission int64 `json:"commission,omitempty" xml:"commission,omitempty"`
	// 未分类商家总优惠，单位：分
	OtherMerchantSubsidyFee int64 `json:"other_merchant_subsidy_fee,omitempty" xml:"other_merchant_subsidy_fee,omitempty"`
	// 未分类平台总优惠，单位：分
	OtherPlatSubsidyFee int64 `json:"other_plat_subsidy_fee,omitempty" xml:"other_plat_subsidy_fee,omitempty"`
	// 商家应退给平台的金额，单位：分
	MerchantTotalFee int64 `json:"merchant_total_fee,omitempty" xml:"merchant_total_fee,omitempty"`
	// 退的渠道配送费补贴，单位：分
	PlatSendSubsidyFee int64 `json:"plat_send_subsidy_fee,omitempty" xml:"plat_send_subsidy_fee,omitempty"`
	// 退的商家配送费补贴，单位：分
	MerchantSendSubsidyFee int64 `json:"merchant_send_subsidy_fee,omitempty" xml:"merchant_send_subsidy_fee,omitempty"`
	// 商家呼单小费，单位：分
	MerchantCallOrderFee int64 `json:"merchant_call_order_fee,omitempty" xml:"merchant_call_order_fee,omitempty"`
	// 冷链配送费，单位：分
	ColdChainSendFee int64 `json:"cold_chain_send_fee,omitempty" xml:"cold_chain_send_fee,omitempty"`
	// 商家呼单配送费，单位：分
	MerchantCallOrderSendFee int64 `json:"merchant_call_order_send_fee,omitempty" xml:"merchant_call_order_send_fee,omitempty"`
	// 配送保险，单位：分
	SendInsuranceFee int64 `json:"send_insurance_fee,omitempty" xml:"send_insurance_fee,omitempty"`
	// 物流驻店服务费，单位：分
	LogisticsShopServiceFee int64 `json:"logistics_shop_service_fee,omitempty" xml:"logistics_shop_service_fee,omitempty"`
	// 实收增值服务费，单位：分
	ActualIncrementServiceFee int64 `json:"actual_increment_service_fee,omitempty" xml:"actual_increment_service_fee,omitempty"`
	// 履约增值服务费，单位：分
	PerformanceIncrementServiceFee int64 `json:"performance_increment_service_fee,omitempty" xml:"performance_increment_service_fee,omitempty"`
	// 距离加价履约费，单位：分
	DistanceIncreasePerformanceFee int64 `json:"distance_increase_performance_fee,omitempty" xml:"distance_increase_performance_fee,omitempty"`
	// 时段加价履约费，单位：分
	TimeIncreasePerformanceFee int64 `json:"time_increase_performance_fee,omitempty" xml:"time_increase_performance_fee,omitempty"`
	// 公益捐款，单位：分
	MerchantPublicDonation int64 `json:"merchant_public_donation,omitempty" xml:"merchant_public_donation,omitempty"`
	// 用户积分抵扣金额，单位：分
	PlatPointsDeductionFee int64 `json:"plat_points_deduction_fee,omitempty" xml:"plat_points_deduction_fee,omitempty"`
	// 自提服务费优惠金额，单位：分
	SelfPickDiscountFee int64 `json:"self_pick_discount_fee,omitempty" xml:"self_pick_discount_fee,omitempty"`
	// 自提服务费，单位：分
	SelfPickPayableFee int64 `json:"self_pick_payable_fee,omitempty" xml:"self_pick_payable_fee,omitempty"`
	// 商家基础配送费，单位：分
	MerchantBaseSendFee int64 `json:"merchant_base_send_fee,omitempty" xml:"merchant_base_send_fee,omitempty"`
	// 价格加价履约费，单位：分
	PriceIncreasePerformanceFee int64 `json:"price_increase_performance_fee,omitempty" xml:"price_increase_performance_fee,omitempty"`
	// 退款原因类型：10-整单全缺;11-订单部分缺;12-商家侧发起的退款(商家拒单); 13-用户侧发起的退款(用户逆向申请);15-用户拒收
	RefundReasonType int64 `json:"refund_reason_type,omitempty" xml:"refund_reason_type,omitempty"`
}

var poolRefundCompleteInfo = sync.Pool{
	New: func() any {
		return new(RefundCompleteInfo)
	},
}

// GetRefundCompleteInfo() 从对象池中获取RefundCompleteInfo
func GetRefundCompleteInfo() *RefundCompleteInfo {
	return poolRefundCompleteInfo.Get().(*RefundCompleteInfo)
}

// ReleaseRefundCompleteInfo 释放RefundCompleteInfo
func ReleaseRefundCompleteInfo(v *RefundCompleteInfo) {
	v.SubRefundOrders = v.SubRefundOrders[:0]
	v.PayChannels = v.PayChannels[:0]
	v.OutOrderId = ""
	v.OutShopId = ""
	v.OutRefundId = ""
	v.ShopId = ""
	v.RefundFee = 0
	v.RefundPostFee = 0
	v.RefundPackageFee = 0
	v.OrderFrom = 0
	v.Commission = 0
	v.OtherMerchantSubsidyFee = 0
	v.OtherPlatSubsidyFee = 0
	v.MerchantTotalFee = 0
	v.PlatSendSubsidyFee = 0
	v.MerchantSendSubsidyFee = 0
	v.MerchantCallOrderFee = 0
	v.ColdChainSendFee = 0
	v.MerchantCallOrderSendFee = 0
	v.SendInsuranceFee = 0
	v.LogisticsShopServiceFee = 0
	v.ActualIncrementServiceFee = 0
	v.PerformanceIncrementServiceFee = 0
	v.DistanceIncreasePerformanceFee = 0
	v.TimeIncreasePerformanceFee = 0
	v.MerchantPublicDonation = 0
	v.PlatPointsDeductionFee = 0
	v.SelfPickDiscountFee = 0
	v.SelfPickPayableFee = 0
	v.MerchantBaseSendFee = 0
	v.PriceIncreasePerformanceFee = 0
	v.RefundReasonType = 0
	poolRefundCompleteInfo.Put(v)
}
