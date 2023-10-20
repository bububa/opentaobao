package tblogistics

import (
	"sync"
)

// CheckDeliveryAuthTopResponse 结构体
type CheckDeliveryAuthTopResponse struct {
	// 扩展
	Features string `json:"features,omitempty" xml:"features,omitempty"`
	// 0正常,1异常
	AuthStatus int64 `json:"auth_status,omitempty" xml:"auth_status,omitempty"`
}

var poolCheckDeliveryAuthTopResponse = sync.Pool{
	New: func() any {
		return new(CheckDeliveryAuthTopResponse)
	},
}

// GetCheckDeliveryAuthTopResponse() 从对象池中获取CheckDeliveryAuthTopResponse
func GetCheckDeliveryAuthTopResponse() *CheckDeliveryAuthTopResponse {
	return poolCheckDeliveryAuthTopResponse.Get().(*CheckDeliveryAuthTopResponse)
}

// ReleaseCheckDeliveryAuthTopResponse 释放CheckDeliveryAuthTopResponse
func ReleaseCheckDeliveryAuthTopResponse(v *CheckDeliveryAuthTopResponse) {
	v.Features = ""
	v.AuthStatus = 0
	poolCheckDeliveryAuthTopResponse.Put(v)
}
