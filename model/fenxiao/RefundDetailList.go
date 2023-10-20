package fenxiao

import (
	"sync"
)

// RefundDetailList 结构体
type RefundDetailList struct {
	// 退货的物流信息
	ReturnLogistics []RefundLogistics `json:"return_logistics,omitempty" xml:"return_logistics>refund_logistics,omitempty"`
	// 退款明细项，记录退款涉及的订单
	RefundItems []RefundItem `json:"refund_items,omitempty" xml:"refund_items>refund_item,omitempty"`
	// 退款单最后修改时间，格式 yyyy-MM-dd HH:mm:ss
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 退款说明
	RefundDesc string `json:"refund_desc,omitempty" xml:"refund_desc,omitempty"`
	// 退款金额，单位：元
	RefundFeeYuan string `json:"refund_fee_yuan,omitempty" xml:"refund_fee_yuan,omitempty"`
	// 支付给供应商的金额，单位：元
	PaySupFeeYuan string `json:"pay_sup_fee_yuan,omitempty" xml:"pay_sup_fee_yuan,omitempty"`
	// 退款创建时间
	RefundCreateTime string `json:"refund_create_time,omitempty" xml:"refund_create_time,omitempty"`
	// 退款单状态code，RF_STATUS_NO_REFUND-已撤回；RF_STATUS_WAIT_SELLER_AGREE-等待卖家同意；RF_STATUS_SELLER_REFUSE-卖家拒绝退款,等待买家修改；RF_STATUS_WAIT_BUYER_RETURN_GOODS-等待买家退货；RF_STATUS_WAIT_SELLER_CONFIRM_GOODS-买家退货,等待卖家确认收货；RF_STATUS_SELLER_REFUSE_RETURN_GOODS-卖家拒绝收货；RF_STATUS_WAIT_SYSTEM_TRANSFER-等待系统打款；RF_STATUS_SUCCESS-退款成功；RF_STATUS_CLOSED-退款关闭
	RefundStatusCode string `json:"refund_status_code,omitempty" xml:"refund_status_code,omitempty"`
	// 退款金额，单位：分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退款单状态，10-已撤回；20-等待卖家同意；30-卖家拒绝退款，等待买家修改；40-等待买家退货；50-买家退货，等待卖家确认收货；60-卖家拒绝收货；90-等待系统打款；100-退款成功；200-退款关闭
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 发生退款的分销子采购单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 买家退款单信息
	BuyerRefund *BuyerRefund `json:"buyer_refund,omitempty" xml:"buyer_refund,omitempty"`
	// 分销采购单id
	PurchaseOrderId int64 `json:"purchase_order_id,omitempty" xml:"purchase_order_id,omitempty"`
	// 支付给供应商的金额，单位：分
	PaySupFee int64 `json:"pay_sup_fee,omitempty" xml:"pay_sup_fee,omitempty"`
	// 退款单类型，10-未发货退款；20-已发货退货；30-已发货仅退款；40-拒收
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	// 退款单id
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 是否退货
	ReturnGoods bool `json:"return_goods,omitempty" xml:"return_goods,omitempty"`
}

var poolRefundDetailList = sync.Pool{
	New: func() any {
		return new(RefundDetailList)
	},
}

// GetRefundDetailList() 从对象池中获取RefundDetailList
func GetRefundDetailList() *RefundDetailList {
	return poolRefundDetailList.Get().(*RefundDetailList)
}

// ReleaseRefundDetailList 释放RefundDetailList
func ReleaseRefundDetailList(v *RefundDetailList) {
	v.ReturnLogistics = v.ReturnLogistics[:0]
	v.RefundItems = v.RefundItems[:0]
	v.GmtModified = ""
	v.RefundDesc = ""
	v.RefundFeeYuan = ""
	v.PaySupFeeYuan = ""
	v.RefundCreateTime = ""
	v.RefundStatusCode = ""
	v.RefundFee = 0
	v.RefundStatus = 0
	v.SubOrderId = 0
	v.BuyerRefund = nil
	v.PurchaseOrderId = 0
	v.PaySupFee = 0
	v.RefundType = 0
	v.RefundId = 0
	v.ReturnGoods = false
	poolRefundDetailList.Put(v)
}
