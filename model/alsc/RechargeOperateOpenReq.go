package alsc

import (
	"sync"
)

// RechargeOperateOpenReq 结构体
type RechargeOperateOpenReq struct {
	// saas品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 顾客id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 操作人id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 外部id类型，wechat：微信openId ；alipay：支付宝
	OuterType string `json:"outer_type,omitempty" xml:"outer_type,omitempty"`
	// 操作备注
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 请求幂等id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// saas门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 回退单号,operate_type为4退款时必填
	NewOuterOrderId string `json:"new_outer_order_id,omitempty" xml:"new_outer_order_id,omitempty"`
	// 操作储值
	ChangeValue int64 `json:"change_value,omitempty" xml:"change_value,omitempty"`
	// 1-充值(charge) 2-冻结(freeze) 3-核销(verify) 4-充值回退(refund)
	OperateType int64 `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
}

var poolRechargeOperateOpenReq = sync.Pool{
	New: func() any {
		return new(RechargeOperateOpenReq)
	},
}

// GetRechargeOperateOpenReq() 从对象池中获取RechargeOperateOpenReq
func GetRechargeOperateOpenReq() *RechargeOperateOpenReq {
	return poolRechargeOperateOpenReq.Get().(*RechargeOperateOpenReq)
}

// ReleaseRechargeOperateOpenReq 释放RechargeOperateOpenReq
func ReleaseRechargeOperateOpenReq(v *RechargeOperateOpenReq) {
	v.BrandId = ""
	v.CustomerId = ""
	v.Mobile = ""
	v.OperatorId = ""
	v.OuterId = ""
	v.OuterOrderId = ""
	v.OuterType = ""
	v.Reason = ""
	v.RequestId = ""
	v.ShopId = ""
	v.NewOuterOrderId = ""
	v.ChangeValue = 0
	v.OperateType = 0
	poolRechargeOperateOpenReq.Put(v)
}
