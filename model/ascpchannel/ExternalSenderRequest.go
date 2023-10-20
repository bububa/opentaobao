package ascpchannel

import (
	"sync"
)

// ExternalSenderRequest 结构体
type ExternalSenderRequest struct {
	// 发货人 手机号
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 发货人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
}

var poolExternalSenderRequest = sync.Pool{
	New: func() any {
		return new(ExternalSenderRequest)
	},
}

// GetExternalSenderRequest() 从对象池中获取ExternalSenderRequest
func GetExternalSenderRequest() *ExternalSenderRequest {
	return poolExternalSenderRequest.Get().(*ExternalSenderRequest)
}

// ReleaseExternalSenderRequest 释放ExternalSenderRequest
func ReleaseExternalSenderRequest(v *ExternalSenderRequest) {
	v.MobilePhone = ""
	v.ContactName = ""
	poolExternalSenderRequest.Put(v)
}
