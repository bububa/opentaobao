package lsttrade

import (
	"sync"
)

// BaseInfo 结构体
type BaseInfo struct {
	// 到达时间
	AllDeliveredTime string `json:"all_delivered_time,omitempty" xml:"all_delivered_time,omitempty"`
	// 备注
	BuyerRemarkIcon string `json:"buyer_remark_icon,omitempty" xml:"buyer_remark_icon,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 供应商名称
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// 订单的售中退款状态，等待卖家同意：waitselleragree ，待买家修改：waitbuyermodify，等待买家退货：waitbuyersend，等待卖家确认收货：waitsellerreceive，退款成功：refundsuccess，退款失败：refundclose
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 外部支付交易Id
	AlipayTradeId string `json:"alipay_trade_id,omitempty" xml:"alipay_trade_id,omitempty"`
	// 卖家loginId
	SellerLoginId string `json:"seller_login_id,omitempty" xml:"seller_login_id,omitempty"`
	// 买家留言，不超过500字
	BuyerFeedback string `json:"buyer_feedback,omitempty" xml:"buyer_feedback,omitempty"`
	// 4.0交易流程模板code
	FlowTemplateCode string `json:"flow_template_code,omitempty" xml:"flow_template_code,omitempty"`
	// 买家loginId，旺旺Id
	BuyerLoginId string `json:"buyer_login_id,omitempty" xml:"buyer_login_id,omitempty"`
	// 修改时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 确认时间
	ConfirmedTime string `json:"confirmed_time,omitempty" xml:"confirmed_time,omitempty"`
	// store_name小店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 关闭原因。buyerCancel:买家取消订单，sellerGoodsLack:卖家库存不足，other:其它
	CloseReason string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// 1:担保交易 2:预存款交易 3:ETC境外收单交易 4:即时到帐交易 5:保障金安全交易 6:统一交易流程 7:分阶段付款 8.货到付款交易 9.信用凭证支付交易 10.账期支付交易，50060 交易4.0
	TradeType string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 订单状态交易状态，waitbuyerpay:等待买家付款;waitsellersend:
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 线下订单业务员id
	OfflineYdId string `json:"offline_yd_id,omitempty" xml:"offline_yd_id,omitempty"`
	// 线下订单业务员name
	OfflineYdName string `json:"offline_yd_name,omitempty" xml:"offline_yd_name,omitempty"`
	// 接收信息
	ReceiverInfo *ReceiverInfo `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
	// 折扣信息-优惠信息，（含优惠券-不包含买家红包）
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 货品金额
	SumProductPayment int64 `json:"sum_product_payment,omitempty" xml:"sum_product_payment,omitempty"`
	// 主订单id
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 卖家联系人信息
	SellerContact *SellerContact `json:"seller_contact,omitempty" xml:"seller_contact,omitempty"`
	// 买家信息
	BuyerContact *BuyerContact `json:"buyer_contact,omitempty" xml:"buyer_contact,omitempty"`
	// 运费
	ShippingFee int64 `json:"shipping_fee,omitempty" xml:"shipping_fee,omitempty"`
	// 实际支付金额
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 小店id
	LeadsId int64 `json:"leads_id,omitempty" xml:"leads_id,omitempty"`
	// 是否线下订单
	OfflineOrder bool `json:"offline_order,omitempty" xml:"offline_order,omitempty"`
	// 是否车销订单
	OfflineCarOrder bool `json:"offline_car_order,omitempty" xml:"offline_car_order,omitempty"`
}

var poolBaseInfo = sync.Pool{
	New: func() any {
		return new(BaseInfo)
	},
}

// GetBaseInfo() 从对象池中获取BaseInfo
func GetBaseInfo() *BaseInfo {
	return poolBaseInfo.Get().(*BaseInfo)
}

// ReleaseBaseInfo 释放BaseInfo
func ReleaseBaseInfo(v *BaseInfo) {
	v.AllDeliveredTime = ""
	v.BuyerRemarkIcon = ""
	v.PayTime = ""
	v.SellerName = ""
	v.RefundStatus = ""
	v.AlipayTradeId = ""
	v.SellerLoginId = ""
	v.BuyerFeedback = ""
	v.FlowTemplateCode = ""
	v.BuyerLoginId = ""
	v.ModifyTime = ""
	v.ConfirmedTime = ""
	v.StoreName = ""
	v.CloseReason = ""
	v.TradeType = ""
	v.OrderStatus = ""
	v.CreateTime = ""
	v.OfflineYdId = ""
	v.OfflineYdName = ""
	v.ReceiverInfo = nil
	v.Discount = 0
	v.SumProductPayment = 0
	v.MainOrderId = 0
	v.SellerContact = nil
	v.BuyerContact = nil
	v.ShippingFee = 0
	v.TotalAmount = 0
	v.LeadsId = 0
	v.OfflineOrder = false
	v.OfflineCarOrder = false
	poolBaseInfo.Put(v)
}
