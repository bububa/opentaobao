package vaccin

import (
	"sync"
)

// CancelVcRegisterRequest 结构体
type CancelVcRegisterRequest struct {
	// cdc侧的登记单id
	RegisterId string `json:"register_id,omitempty" xml:"register_id,omitempty"`
	// 健康侧的用户支付宝id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 取消登记时间，会校验格式 yyyy-MM-dd HH:mm:ss
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
}

var poolCancelVcRegisterRequest = sync.Pool{
	New: func() any {
		return new(CancelVcRegisterRequest)
	},
}

// GetCancelVcRegisterRequest() 从对象池中获取CancelVcRegisterRequest
func GetCancelVcRegisterRequest() *CancelVcRegisterRequest {
	return poolCancelVcRegisterRequest.Get().(*CancelVcRegisterRequest)
}

// ReleaseCancelVcRegisterRequest 释放CancelVcRegisterRequest
func ReleaseCancelVcRegisterRequest(v *CancelVcRegisterRequest) {
	v.RegisterId = ""
	v.UserId = ""
	v.CancelTime = ""
	poolCancelVcRegisterRequest.Put(v)
}
