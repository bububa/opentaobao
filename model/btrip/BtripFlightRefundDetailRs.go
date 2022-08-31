package btrip

import (
	"github.com/bububa/opentaobao/model"
)

// BtripFlightRefundDetailRs 结构体
type BtripFlightRefundDetailRs struct {
	// 退票费用列表
	RefundFeeList []RefundFeeInfo `json:"refund_fee_list,omitempty" xml:"refund_fee_list>refund_fee_info,omitempty"`
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 分销外部退票单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 退票原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商旅订单号
	BtripOrderId int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// 商旅退票单号
	BtripSubOrderId int64 `json:"btrip_sub_order_id,omitempty" xml:"btrip_sub_order_id,omitempty"`
	// 是否是自愿退票
	IsVoluntary *model.File `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	// 退票金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退票金额
	RefundPrice int64 `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
}
