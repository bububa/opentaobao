package ascp

import (
	"sync"
)

// AddressBlacklistDeleteResponse 结构体
type AddressBlacklistDeleteResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否可重试
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
}

var poolAddressBlacklistDeleteResponse = sync.Pool{
	New: func() any {
		return new(AddressBlacklistDeleteResponse)
	},
}

// GetAddressBlacklistDeleteResponse() 从对象池中获取AddressBlacklistDeleteResponse
func GetAddressBlacklistDeleteResponse() *AddressBlacklistDeleteResponse {
	return poolAddressBlacklistDeleteResponse.Get().(*AddressBlacklistDeleteResponse)
}

// ReleaseAddressBlacklistDeleteResponse 释放AddressBlacklistDeleteResponse
func ReleaseAddressBlacklistDeleteResponse(v *AddressBlacklistDeleteResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.IsRetry = false
	poolAddressBlacklistDeleteResponse.Put(v)
}
