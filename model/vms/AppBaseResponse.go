package vms

import (
	"sync"
)

// AppBaseResponse 结构体
type AppBaseResponse struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAppBaseResponse = sync.Pool{
	New: func() any {
		return new(AppBaseResponse)
	},
}

// GetAppBaseResponse() 从对象池中获取AppBaseResponse
func GetAppBaseResponse() *AppBaseResponse {
	return poolAppBaseResponse.Get().(*AppBaseResponse)
}

// ReleaseAppBaseResponse 释放AppBaseResponse
func ReleaseAppBaseResponse(v *AppBaseResponse) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = false
	v.Success = false
	poolAppBaseResponse.Put(v)
}
