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
	// 重试时给用户提示的文案
	RetryClientTips string `json:"retry_client_tips,omitempty" xml:"retry_client_tips,omitempty"`
	// 改签费用
	ChangeFee int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// 改签单状态
	Status *model.File `json:"status,omitempty" xml:"status,omitempty"`
	// 升舱费
	UpgradeFee int64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
	// 商旅改签申请单号
	BtripSubOrderId int64 `json:"btrip_sub_order_id,omitempty" xml:"btrip_sub_order_id,omitempty"`
	// 最大的重试次数
	MaxRetryTimes int64 `json:"max_retry_times,omitempty" xml:"max_retry_times,omitempty"`
	// 重试的间隔时间单位ms
	NextRetryInterval int64 `json:"next_retry_interval,omitempty" xml:"next_retry_interval,omitempty"`
	// 预定后变价-用户原支付费用总金额
	BookingChangedTotalFee int64 `json:"booking_changed_total_fee,omitempty" xml:"booking_changed_total_fee,omitempty"`
	// 预定后变价-用户原支付费用总金额
	BookingOriginTotalFee int64 `json:"booking_origin_total_fee,omitempty" xml:"booking_origin_total_fee,omitempty"`
	// 能否支付
	CanPay bool `json:"can_pay,omitempty" xml:"can_pay,omitempty"`
	// 是否重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
	// 是否变价
	BookingPriceChanged bool `json:"booking_price_changed,omitempty" xml:"booking_price_changed,omitempty"`
}
