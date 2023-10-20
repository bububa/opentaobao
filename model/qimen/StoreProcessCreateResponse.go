package qimen

import (
	"sync"
)

// StoreProcessCreateResponse 结构体
type StoreProcessCreateResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 仓储系统处理单ID
	ProcessOrderId string `json:"processOrderId,omitempty" xml:"processOrderId,omitempty"`
}

var poolStoreProcessCreateResponse = sync.Pool{
	New: func() any {
		return new(StoreProcessCreateResponse)
	},
}

// GetStoreProcessCreateResponse() 从对象池中获取StoreProcessCreateResponse
func GetStoreProcessCreateResponse() *StoreProcessCreateResponse {
	return poolStoreProcessCreateResponse.Get().(*StoreProcessCreateResponse)
}

// ReleaseStoreProcessCreateResponse 释放StoreProcessCreateResponse
func ReleaseStoreProcessCreateResponse(v *StoreProcessCreateResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.ProcessOrderId = ""
	poolStoreProcessCreateResponse.Put(v)
}
