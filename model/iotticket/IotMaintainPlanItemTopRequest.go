package iotticket

import (
	"sync"
)

// IotMaintainPlanItemTopRequest 结构体
type IotMaintainPlanItemTopRequest struct {
	// 付款角色：merchant-商家记账；customer-客户支付
	PayRole string `json:"pay_role,omitempty" xml:"pay_role,omitempty"`
	// 设备编码（需要映射）
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
}

var poolIotMaintainPlanItemTopRequest = sync.Pool{
	New: func() any {
		return new(IotMaintainPlanItemTopRequest)
	},
}

// GetIotMaintainPlanItemTopRequest() 从对象池中获取IotMaintainPlanItemTopRequest
func GetIotMaintainPlanItemTopRequest() *IotMaintainPlanItemTopRequest {
	return poolIotMaintainPlanItemTopRequest.Get().(*IotMaintainPlanItemTopRequest)
}

// ReleaseIotMaintainPlanItemTopRequest 释放IotMaintainPlanItemTopRequest
func ReleaseIotMaintainPlanItemTopRequest(v *IotMaintainPlanItemTopRequest) {
	v.PayRole = ""
	v.ItemCode = ""
	poolIotMaintainPlanItemTopRequest.Put(v)
}
