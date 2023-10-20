package tmallhk

import (
	"sync"
)

// OrderClearanceInfoRequest 结构体
type OrderClearanceInfoRequest struct {
	// 清关结果编码放行/000开头 清关失败/100开头
	ClearanceResultCode string `json:"clearance_result_code,omitempty" xml:"clearance_result_code,omitempty"`
	// 清关口岸代码
	ClearancePortNo string `json:"clearance_port_no,omitempty" xml:"clearance_port_no,omitempty"`
	// 业务清关触发时间
	ClearanceTime string `json:"clearance_time,omitempty" xml:"clearance_time,omitempty"`
	// 海关放行的清单编号
	CustomsPassNo string `json:"customs_pass_no,omitempty" xml:"customs_pass_no,omitempty"`
	// 海关回执
	CustomsReturnReceipt string `json:"customs_return_receipt,omitempty" xml:"customs_return_receipt,omitempty"`
	// 实名信息幂等健，来源于清关材料接口
	Idempotent string `json:"idempotent,omitempty" xml:"idempotent,omitempty"`
	// 清关状态 1清关发起中回调  2 清关结果回调
	ClearanceStatus int64 `json:"clearance_status,omitempty" xml:"clearance_status,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolOrderClearanceInfoRequest = sync.Pool{
	New: func() any {
		return new(OrderClearanceInfoRequest)
	},
}

// GetOrderClearanceInfoRequest() 从对象池中获取OrderClearanceInfoRequest
func GetOrderClearanceInfoRequest() *OrderClearanceInfoRequest {
	return poolOrderClearanceInfoRequest.Get().(*OrderClearanceInfoRequest)
}

// ReleaseOrderClearanceInfoRequest 释放OrderClearanceInfoRequest
func ReleaseOrderClearanceInfoRequest(v *OrderClearanceInfoRequest) {
	v.ClearanceResultCode = ""
	v.ClearancePortNo = ""
	v.ClearanceTime = ""
	v.CustomsPassNo = ""
	v.CustomsReturnReceipt = ""
	v.Idempotent = ""
	v.ClearanceStatus = 0
	v.OrderId = 0
	poolOrderClearanceInfoRequest.Put(v)
}
