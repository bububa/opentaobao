package ascp

import (
	"sync"
)

// AddressBlacklistResponse 结构体
type AddressBlacklistResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否可重试
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
}

var poolAddressBlacklistResponse = sync.Pool{
	New: func() any {
		return new(AddressBlacklistResponse)
	},
}

// GetAddressBlacklistResponse() 从对象池中获取AddressBlacklistResponse
func GetAddressBlacklistResponse() *AddressBlacklistResponse {
	return poolAddressBlacklistResponse.Get().(*AddressBlacklistResponse)
}

// ReleaseAddressBlacklistResponse 释放AddressBlacklistResponse
func ReleaseAddressBlacklistResponse(v *AddressBlacklistResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.IsRetry = false
	poolAddressBlacklistResponse.Put(v)
}
