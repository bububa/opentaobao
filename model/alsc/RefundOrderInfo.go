package alsc

import (
	"sync"
)

// RefundOrderInfo 结构体
type RefundOrderInfo struct {
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 设备ip
	DeviceIp string `json:"device_ip,omitempty" xml:"device_ip,omitempty"`
	// 授权人id
	OutAuthorizerId string `json:"out_authorizer_id,omitempty" xml:"out_authorizer_id,omitempty"`
	// 授权人名称
	OutAuthorizerName string `json:"out_authorizer_name,omitempty" xml:"out_authorizer_name,omitempty"`
	// 操作人id
	OutOperatorId string `json:"out_operator_id,omitempty" xml:"out_operator_id,omitempty"`
	// 操作人名称
	OutOperatorName string `json:"out_operator_name,omitempty" xml:"out_operator_name,omitempty"`
	// 外部订单号
	OutOrderNo string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	// 业务方退款单号
	OutRefundOrderNo string `json:"out_refund_order_no,omitempty" xml:"out_refund_order_no,omitempty"`
	// 退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 退款类型 REFUND_ITEM 退菜退款 ONLY_REFUND仅退款款ANTI_SEELEMENT 反结账
	RefundType string `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	// 退款单状态 SUCCESS 成功；  FAIL失败
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退款成功时间
	SuccessTime int64 `json:"success_time,omitempty" xml:"success_time,omitempty"`
}

var poolRefundOrderInfo = sync.Pool{
	New: func() any {
		return new(RefundOrderInfo)
	},
}

// GetRefundOrderInfo() 从对象池中获取RefundOrderInfo
func GetRefundOrderInfo() *RefundOrderInfo {
	return poolRefundOrderInfo.Get().(*RefundOrderInfo)
}

// ReleaseRefundOrderInfo 释放RefundOrderInfo
func ReleaseRefundOrderInfo(v *RefundOrderInfo) {
	v.DeviceId = ""
	v.DeviceIp = ""
	v.OutAuthorizerId = ""
	v.OutAuthorizerName = ""
	v.OutOperatorId = ""
	v.OutOperatorName = ""
	v.OutOrderNo = ""
	v.OutRefundOrderNo = ""
	v.RefundReason = ""
	v.RefundType = ""
	v.Status = ""
	v.RefundFee = 0
	v.SuccessTime = 0
	poolRefundOrderInfo.Put(v)
}
