package wdk

import (
	"sync"
)

// OrderInfo 结构体
type OrderInfo struct {
	// 子订单信息
	SubOrders []SubOrder `json:"sub_orders,omitempty" xml:"sub_orders>sub_order,omitempty"`
	// 支付渠道，不填会默认使用类型1，支付金额=主单的payFee
	PayChannels []PayChannel `json:"pay_channels,omitempty" xml:"pay_channels>pay_channel,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 外部渠道店ID(与shop_id必选其一)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 创单时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 订单小号
	SerialNo string `json:"serial_no,omitempty" xml:"serial_no,omitempty"`
	// 渠道店id(与out_shop_id必选其一)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 订单扩展数据
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 经营店Id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 买家信息
	Buyer *Buyer `json:"buyer,omitempty" xml:"buyer,omitempty"`
	// 收件人信息
	Consignee *Consignee `json:"consignee,omitempty" xml:"consignee,omitempty"`
	// 实际支付金额
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 原始金额
	OriginFee int64 `json:"origin_fee,omitempty" xml:"origin_fee,omitempty"`
	// 优惠金额
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 运费
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 订单来源(选填out_shop_id时该值必填)
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 配送方式 1:平台配送 2:商家自配送 3:到店自提
	PickupType int64 `json:"pickup_type,omitempty" xml:"pickup_type,omitempty"`
	// 平台佣金
	Commission int64 `json:"commission,omitempty" xml:"commission,omitempty"`
	// 包装费
	PackageFee int64 `json:"package_fee,omitempty" xml:"package_fee,omitempty"`
	// 商家应收总金额
	MerchantTotalFee int64 `json:"merchant_total_fee,omitempty" xml:"merchant_total_fee,omitempty"`
	// 未分类商家总优惠
	OtherMerchantSubsidyFee int64 `json:"other_merchant_subsidy_fee,omitempty" xml:"other_merchant_subsidy_fee,omitempty"`
	// 未分类平台总优惠
	OtherPlatSubsidyFee int64 `json:"other_plat_subsidy_fee,omitempty" xml:"other_plat_subsidy_fee,omitempty"`
	// 商家基础配送费
	MerchantBaseSendFee int64 `json:"merchant_base_send_fee,omitempty" xml:"merchant_base_send_fee,omitempty"`
	// 平台配送费补贴
	PlatSendSubsidyFee int64 `json:"plat_send_subsidy_fee,omitempty" xml:"plat_send_subsidy_fee,omitempty"`
	// 商家配送费补贴
	MerchantSendSubsidyFee int64 `json:"merchant_send_subsidy_fee,omitempty" xml:"merchant_send_subsidy_fee,omitempty"`
	// 商家呼单小费
	MerchantCallOrderFee int64 `json:"merchant_call_order_fee,omitempty" xml:"merchant_call_order_fee,omitempty"`
	// 冷链配送费
	ColdChainSendFee int64 `json:"cold_chain_send_fee,omitempty" xml:"cold_chain_send_fee,omitempty"`
	// 商家呼单配送费
	MerchantCallOrderSendFee int64 `json:"merchant_call_order_send_fee,omitempty" xml:"merchant_call_order_send_fee,omitempty"`
	// 配送保险
	SendInsuranceFee int64 `json:"send_insurance_fee,omitempty" xml:"send_insurance_fee,omitempty"`
	// 物流驻店服务费
	LogisticsShopServiceFee int64 `json:"logistics_shop_service_fee,omitempty" xml:"logistics_shop_service_fee,omitempty"`
	// 实收增值服务费
	ActualIncrementServiceFee int64 `json:"actual_increment_service_fee,omitempty" xml:"actual_increment_service_fee,omitempty"`
	// 履约增值服务费
	PerformanceIncrementServiceFee int64 `json:"performance_increment_service_fee,omitempty" xml:"performance_increment_service_fee,omitempty"`
	// 距离加价履约费
	DistanceIncreasePerformanceFee int64 `json:"distance_increase_performance_fee,omitempty" xml:"distance_increase_performance_fee,omitempty"`
	// 时段加价履约费
	TimeIncreasePerformanceFee int64 `json:"time_increase_performance_fee,omitempty" xml:"time_increase_performance_fee,omitempty"`
	// 公益捐款
	MerchantPublicDonation int64 `json:"merchant_public_donation,omitempty" xml:"merchant_public_donation,omitempty"`
	// 用户积分抵扣金额
	PlatPointsDeductionFee int64 `json:"plat_points_deduction_fee,omitempty" xml:"plat_points_deduction_fee,omitempty"`
	// 自提服务费优惠金额
	SelfPickDiscountFee int64 `json:"self_pick_discount_fee,omitempty" xml:"self_pick_discount_fee,omitempty"`
	// 自提服务费
	SelfPickPayableFee int64 `json:"self_pick_payable_fee,omitempty" xml:"self_pick_payable_fee,omitempty"`
	// 价格加价履约费
	PriceIncreasePerformanceFee int64 `json:"price_increase_performance_fee,omitempty" xml:"price_increase_performance_fee,omitempty"`
}

var poolOrderInfo = sync.Pool{
	New: func() any {
		return new(OrderInfo)
	},
}

// GetOrderInfo() 从对象池中获取OrderInfo
func GetOrderInfo() *OrderInfo {
	return poolOrderInfo.Get().(*OrderInfo)
}

// ReleaseOrderInfo 释放OrderInfo
func ReleaseOrderInfo(v *OrderInfo) {
	v.SubOrders = v.SubOrders[:0]
	v.PayChannels = v.PayChannels[:0]
	v.OutOrderId = ""
	v.OutShopId = ""
	v.OrderStatus = ""
	v.PayTime = ""
	v.CreateTime = ""
	v.SerialNo = ""
	v.ShopId = ""
	v.Ext = ""
	v.StoreId = ""
	v.Buyer = nil
	v.Consignee = nil
	v.PayFee = 0
	v.OriginFee = 0
	v.DiscountFee = 0
	v.PostFee = 0
	v.OrderFrom = 0
	v.PickupType = 0
	v.Commission = 0
	v.PackageFee = 0
	v.MerchantTotalFee = 0
	v.OtherMerchantSubsidyFee = 0
	v.OtherPlatSubsidyFee = 0
	v.MerchantBaseSendFee = 0
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
	v.PriceIncreasePerformanceFee = 0
	poolOrderInfo.Put(v)
}
