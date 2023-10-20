package wdk

import (
	"sync"
)

// OrderResponse 结构体
type OrderResponse struct {
	// 子订单列表
	SubOrderResponseList []SubOrderResponse `json:"sub_order_response_list,omitempty" xml:"sub_order_response_list>sub_order_response,omitempty"`
	// 支付渠道信息列表
	PayChannels []OrderPayChannel `json:"pay_channels,omitempty" xml:"pay_channels>order_pay_channel,omitempty"`
	// 主单活动列表
	Activitys []OrderActivity `json:"activitys,omitempty" xml:"activitys>order_activity,omitempty"`
	// 资金优惠
	FundsDiscounts []OrderFundsDiscount `json:"funds_discounts,omitempty" xml:"funds_discounts>order_funds_discount,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 经营店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道店编码
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 渠道订单编码
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 订单状态 PAID = 订单支付完成 PACKAGED = 订单打包出库 SHIPPING = 订单配送揽收 SUCCESS = 交易完成 CLOSE = 订单取消
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 对接渠道的门店ID
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 渠道订单小号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 买家openId
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 订单渠道, 31=淘鲜达等
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 订单编码
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 订单总价，分
	OriginalFee int64 `json:"original_fee,omitempty" xml:"original_fee,omitempty"`
	// 订单总优惠，分
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 配送费，分
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 包装费，分
	PackageFee int64 `json:"package_fee,omitempty" xml:"package_fee,omitempty"`
	// 用户实际支付费用，分
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 收货人信息
	ReceiveInfo *ReceiveInfo `json:"receive_info,omitempty" xml:"receive_info,omitempty"`
	// 订单扩展信息
	Ext *OrderInfoExt `json:"ext,omitempty" xml:"ext,omitempty"`
	// 商品平台优惠分摊
	SkuDiscountPlatformFee int64 `json:"sku_discount_platform_fee,omitempty" xml:"sku_discount_platform_fee,omitempty"`
	// 商品商家优惠分摊
	SkuDiscountMerchantFee int64 `json:"sku_discount_merchant_fee,omitempty" xml:"sku_discount_merchant_fee,omitempty"`
	// 配送费平台优惠分摊�
	PostDiscountPlatformFee int64 `json:"post_discount_platform_fee,omitempty" xml:"post_discount_platform_fee,omitempty"`
	// 配送费商家优惠分摊
	PostDiscountMerchantFee int64 `json:"post_discount_merchant_fee,omitempty" xml:"post_discount_merchant_fee,omitempty"`
	// 配送方式 1=定时配送/ 3=用户自提
	DeliveryType int64 `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
	// 平台扣费对象
	PlatformDeduction *PlatformDeduction `json:"platform_deduction,omitempty" xml:"platform_deduction,omitempty"`
	// 送货信息
	DeliveryInfo *DeliveryInfo `json:"delivery_info,omitempty" xml:"delivery_info,omitempty"`
}

var poolOrderResponse = sync.Pool{
	New: func() any {
		return new(OrderResponse)
	},
}

// GetOrderResponse() 从对象池中获取OrderResponse
func GetOrderResponse() *OrderResponse {
	return poolOrderResponse.Get().(*OrderResponse)
}

// ReleaseOrderResponse 释放OrderResponse
func ReleaseOrderResponse(v *OrderResponse) {
	v.SubOrderResponseList = v.SubOrderResponseList[:0]
	v.PayChannels = v.PayChannels[:0]
	v.Activitys = v.Activitys[:0]
	v.FundsDiscounts = v.FundsDiscounts[:0]
	v.MerchantCode = ""
	v.StoreId = ""
	v.ShopId = ""
	v.OutOrderId = ""
	v.PayTime = ""
	v.OrderStatus = ""
	v.OutShopId = ""
	v.OrderNo = ""
	v.OpenUid = ""
	v.OrderFrom = 0
	v.BizOrderId = 0
	v.OriginalFee = 0
	v.DiscountFee = 0
	v.PostFee = 0
	v.PackageFee = 0
	v.PayFee = 0
	v.ReceiveInfo = nil
	v.Ext = nil
	v.SkuDiscountPlatformFee = 0
	v.SkuDiscountMerchantFee = 0
	v.PostDiscountPlatformFee = 0
	v.PostDiscountMerchantFee = 0
	v.DeliveryType = 0
	v.PlatformDeduction = nil
	v.DeliveryInfo = nil
	poolOrderResponse.Put(v)
}
