package logistic

import (
	"sync"
)

// TmsPickCodeRequest 结构体
type TmsPickCodeRequest struct {
	// 订单业务类型（1-退货业务 ）
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 单据标识
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 订单所需服务，以英文逗号分隔，当前服务枚举值:（1-上门取退
	ServiceFlag string `json:"service_flag,omitempty" xml:"service_flag,omitempty"`
	// 取件码
	PickCode string `json:"pick_code,omitempty" xml:"pick_code,omitempty"`
}

var poolTmsPickCodeRequest = sync.Pool{
	New: func() any {
		return new(TmsPickCodeRequest)
	},
}

// GetTmsPickCodeRequest() 从对象池中获取TmsPickCodeRequest
func GetTmsPickCodeRequest() *TmsPickCodeRequest {
	return poolTmsPickCodeRequest.Get().(*TmsPickCodeRequest)
}

// ReleaseTmsPickCodeRequest 释放TmsPickCodeRequest
func ReleaseTmsPickCodeRequest(v *TmsPickCodeRequest) {
	v.ServiceType = ""
	v.BizCode = ""
	v.ServiceFlag = ""
	v.PickCode = ""
	poolTmsPickCodeRequest.Put(v)
}
