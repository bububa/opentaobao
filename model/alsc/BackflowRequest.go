package alsc

import (
	"sync"
)

// BackflowRequest 结构体
type BackflowRequest struct {
	// 订单或子订单属性信息
	OrderAttributeInfoList []OrderAttributeInfo `json:"order_attribute_info_list,omitempty" xml:"order_attribute_info_list>order_attribute_info,omitempty"`
	// 支付流水明细
	PayDetailInfoList []PayDetailInfo `json:"pay_detail_info_list,omitempty" xml:"pay_detail_info_list>pay_detail_info,omitempty"`
	// 优惠明细
	PromoDetailInfoList []PromoDetailInfo `json:"promo_detail_info_list,omitempty" xml:"promo_detail_info_list>promo_detail_info,omitempty"`
	// 退款资金流水
	RefundFundDetailInfoList []RefundFundDetailInfo `json:"refund_fund_detail_info_list,omitempty" xml:"refund_fund_detail_info_list>refund_fund_detail_info,omitempty"`
	// 退款商品明细
	RefundItemDetailInfoList []RefundItemDetailInfo `json:"refund_item_detail_info_list,omitempty" xml:"refund_item_detail_info_list>refund_item_detail_info,omitempty"`
	// 退款单数据
	RefundOrderInfoList []RefundOrderInfo `json:"refund_order_info_list,omitempty" xml:"refund_order_info_list>refund_order_info,omitempty"`
	// 子订单（商品）信息
	SubOrderInfoList []SubOrderInfo `json:"sub_order_info_list,omitempty" xml:"sub_order_info_list>sub_order_info,omitempty"`
	// 订单来源 KERUYUN  CHENSEN  KPOS
	BizSource string `json:"biz_source,omitempty" xml:"biz_source,omitempty"`
	// 订单信息
	OrderInfo *OrderInfo `json:"order_info,omitempty" xml:"order_info,omitempty"`
}

var poolBackflowRequest = sync.Pool{
	New: func() any {
		return new(BackflowRequest)
	},
}

// GetBackflowRequest() 从对象池中获取BackflowRequest
func GetBackflowRequest() *BackflowRequest {
	return poolBackflowRequest.Get().(*BackflowRequest)
}

// ReleaseBackflowRequest 释放BackflowRequest
func ReleaseBackflowRequest(v *BackflowRequest) {
	v.OrderAttributeInfoList = v.OrderAttributeInfoList[:0]
	v.PayDetailInfoList = v.PayDetailInfoList[:0]
	v.PromoDetailInfoList = v.PromoDetailInfoList[:0]
	v.RefundFundDetailInfoList = v.RefundFundDetailInfoList[:0]
	v.RefundItemDetailInfoList = v.RefundItemDetailInfoList[:0]
	v.RefundOrderInfoList = v.RefundOrderInfoList[:0]
	v.SubOrderInfoList = v.SubOrderInfoList[:0]
	v.BizSource = ""
	v.OrderInfo = nil
	poolBackflowRequest.Put(v)
}
