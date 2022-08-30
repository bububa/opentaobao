package btrip

import (
	"github.com/bububa/opentaobao/model"
)

// BtripFlightModifyApplyRs 结构体
type BtripFlightModifyApplyRs struct {
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 分销外部改签单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 支付截止时间
	DeadlineTime string `json:"deadline_time,omitempty" xml:"deadline_time,omitempty"`
	// 改签费用
	ChangeFee int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// 改签单状态
	Status *model.File `json:"status,omitempty" xml:"status,omitempty"`
	// 升舱费
	UpgradeFee int64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
	// 商旅改签申请单号
	BtripSubOrderId int64 `json:"btrip_sub_order_id,omitempty" xml:"btrip_sub_order_id,omitempty"`
	// 能否支付
	CanPay bool `json:"can_pay,omitempty" xml:"can_pay,omitempty"`
}
