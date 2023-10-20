package alsc

import (
	"sync"
)

// PointOperateOpenReq 结构体
type PointOperateOpenReq struct {
	// saas品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 操作人id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部id类型，wechat：微信openId  alipay：支付宝
	OuterType string `json:"outer_type,omitempty" xml:"outer_type,omitempty"`
	// 操作原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 请求幂等id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// saas门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 顾客id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 操作积分
	ChangePoint int64 `json:"change_point,omitempty" xml:"change_point,omitempty"`
	// 1-增加(charge)  2-冻结(freeze)  3-核销(verify)  4-扣减(decrease)
	OperateType int64 `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
}

var poolPointOperateOpenReq = sync.Pool{
	New: func() any {
		return new(PointOperateOpenReq)
	},
}

// GetPointOperateOpenReq() 从对象池中获取PointOperateOpenReq
func GetPointOperateOpenReq() *PointOperateOpenReq {
	return poolPointOperateOpenReq.Get().(*PointOperateOpenReq)
}

// ReleasePointOperateOpenReq 释放PointOperateOpenReq
func ReleasePointOperateOpenReq(v *PointOperateOpenReq) {
	v.BrandId = ""
	v.Mobile = ""
	v.OperatorId = ""
	v.OuterId = ""
	v.OuterType = ""
	v.Reason = ""
	v.RequestId = ""
	v.ShopId = ""
	v.OuterOrderId = ""
	v.CustomerId = ""
	v.ChangePoint = 0
	v.OperateType = 0
	poolPointOperateOpenReq.Put(v)
}
