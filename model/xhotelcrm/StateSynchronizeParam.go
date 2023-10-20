package xhotelcrm

import (
	"sync"
)

// StateSynchronizeParam 结构体
type StateSynchronizeParam struct {
	// 券信息
	CouponActiveParamList []CouponActiveParam `json:"coupon_active_param_list,omitempty" xml:"coupon_active_param_list>coupon_active_param,omitempty"`
	// 卡有效期截止时间
	CardEndDate string `json:"card_end_date,omitempty" xml:"card_end_date,omitempty"`
	// 卡有效期开始时间
	CardStartDate string `json:"card_start_date,omitempty" xml:"card_start_date,omitempty"`
	// action类型
	OpType string `json:"op_type,omitempty" xml:"op_type,omitempty"`
	// action渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 券实例ID
	CouponInstanceId string `json:"coupon_instance_id,omitempty" xml:"coupon_instance_id,omitempty"`
	// 飞猪订单号
	FliggyOrderId string `json:"fliggy_order_id,omitempty" xml:"fliggy_order_id,omitempty"`
	// 卡实例ID
	CardInstanceId string `json:"card_instance_id,omitempty" xml:"card_instance_id,omitempty"`
	// 会员卡号
	MemberCardNo string `json:"member_card_no,omitempty" xml:"member_card_no,omitempty"`
}

var poolStateSynchronizeParam = sync.Pool{
	New: func() any {
		return new(StateSynchronizeParam)
	},
}

// GetStateSynchronizeParam() 从对象池中获取StateSynchronizeParam
func GetStateSynchronizeParam() *StateSynchronizeParam {
	return poolStateSynchronizeParam.Get().(*StateSynchronizeParam)
}

// ReleaseStateSynchronizeParam 释放StateSynchronizeParam
func ReleaseStateSynchronizeParam(v *StateSynchronizeParam) {
	v.CouponActiveParamList = v.CouponActiveParamList[:0]
	v.CardEndDate = ""
	v.CardStartDate = ""
	v.OpType = ""
	v.Channel = ""
	v.CouponInstanceId = ""
	v.FliggyOrderId = ""
	v.CardInstanceId = ""
	v.MemberCardNo = ""
	poolStateSynchronizeParam.Put(v)
}
