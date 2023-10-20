package xhotelonlineorder

import (
	"sync"
)

// OrderNotifyQuery 结构体
type OrderNotifyQuery struct {
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 会员卡号
	MemberCardNo string `json:"member_card_no,omitempty" xml:"member_card_no,omitempty"`
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 通知类型，0代表支付通知
	NotifyType int64 `json:"notify_type,omitempty" xml:"notify_type,omitempty"`
	// 平台促销
	PlatformPromotion int64 `json:"platform_promotion,omitempty" xml:"platform_promotion,omitempty"`
}

var poolOrderNotifyQuery = sync.Pool{
	New: func() any {
		return new(OrderNotifyQuery)
	},
}

// GetOrderNotifyQuery() 从对象池中获取OrderNotifyQuery
func GetOrderNotifyQuery() *OrderNotifyQuery {
	return poolOrderNotifyQuery.Get().(*OrderNotifyQuery)
}

// ReleaseOrderNotifyQuery 释放OrderNotifyQuery
func ReleaseOrderNotifyQuery(v *OrderNotifyQuery) {
	v.PayTime = ""
	v.MemberCardNo = ""
	v.BizOrderId = 0
	v.NotifyType = 0
	v.PlatformPromotion = 0
	poolOrderNotifyQuery.Put(v)
}
