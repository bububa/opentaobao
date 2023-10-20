package wdk

import (
	"sync"
)

// WdkOrderSyncBo 结构体
type WdkOrderSyncBo struct {
	// 子单列表
	SubOrders []Suborders `json:"sub_orders,omitempty" xml:"sub_orders>suborders,omitempty"`
	// 订单优惠信息
	Promotions []Promotions `json:"promotions,omitempty" xml:"promotions>promotions,omitempty"`
	// 外部订单ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 商户编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 买家open_uid
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 期望送达时间段
	ExpectArriveTime string `json:"expect_arrive_time,omitempty" xml:"expect_arrive_time,omitempty"`
	// 订单小号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 五道口订单ID
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 订单来源, 如 TAOBAO (4, &#34;TC自营渠道&#34;),
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 用户订单支付金额,分
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 订单原价,分
	OriginFee int64 `json:"origin_fee,omitempty" xml:"origin_fee,omitempty"`
	// 订单优惠金额,分
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 订单配送费,分
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 订单状态,如PAID_DONE(2, &#34;已付款&#34;), TRADE_SUCCESS(6, &#34;交易成功&#34;)
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 配送方式, 如InTime(1, &#34;即时达&#34;), SetTime(2, &#34;定时达&#34;),TopSpeed(3,&#34;极速达&#34;),NoNeedSend(4,&#34;无需配送&#34;)
	ArriveType int64 `json:"arrive_type,omitempty" xml:"arrive_type,omitempty"`
	// 商家优惠分摊
	DiscountMerchantFee int64 `json:"discount_merchant_fee,omitempty" xml:"discount_merchant_fee,omitempty"`
	// 平台优惠分摊
	DiscountPlatformFee int64 `json:"discount_platform_fee,omitempty" xml:"discount_platform_fee,omitempty"`
	// 包装费
	PackageFee int64 `json:"package_fee,omitempty" xml:"package_fee,omitempty"`
}

var poolWdkOrderSyncBo = sync.Pool{
	New: func() any {
		return new(WdkOrderSyncBo)
	},
}

// GetWdkOrderSyncBo() 从对象池中获取WdkOrderSyncBo
func GetWdkOrderSyncBo() *WdkOrderSyncBo {
	return poolWdkOrderSyncBo.Get().(*WdkOrderSyncBo)
}

// ReleaseWdkOrderSyncBo 释放WdkOrderSyncBo
func ReleaseWdkOrderSyncBo(v *WdkOrderSyncBo) {
	v.SubOrders = v.SubOrders[:0]
	v.Promotions = v.Promotions[:0]
	v.OutOrderId = ""
	v.StoreId = ""
	v.ShopId = ""
	v.MerchantCode = ""
	v.OpenUid = ""
	v.PayTime = ""
	v.ExpectArriveTime = ""
	v.OrderNo = ""
	v.BizOrderId = 0
	v.OrderFrom = 0
	v.PayFee = 0
	v.OriginFee = 0
	v.DiscountFee = 0
	v.PostFee = 0
	v.OrderStatus = 0
	v.ArriveType = 0
	v.DiscountMerchantFee = 0
	v.DiscountPlatformFee = 0
	v.PackageFee = 0
	poolWdkOrderSyncBo.Put(v)
}
