package pur

import (
	"sync"
)

// MallReceivePrResponse 结构体
type MallReceivePrResponse struct {
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误提示
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 返回数据
	Data *MallReceivePrResponseData `json:"data,omitempty" xml:"data,omitempty"`
	// 返回标识
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMallReceivePrResponse = sync.Pool{
	New: func() any {
		return new(MallReceivePrResponse)
	},
}

// GetMallReceivePrResponse() 从对象池中获取MallReceivePrResponse
func GetMallReceivePrResponse() *MallReceivePrResponse {
	return poolMallReceivePrResponse.Get().(*MallReceivePrResponse)
}

// ReleaseMallReceivePrResponse 释放MallReceivePrResponse
func ReleaseMallReceivePrResponse(v *MallReceivePrResponse) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Data = nil
	v.Success = false
	poolMallReceivePrResponse.Put(v)
}
