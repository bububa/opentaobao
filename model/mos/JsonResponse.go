package mos

import (
	"sync"
)

// JsonResponse 结构体
type JsonResponse struct {
	// 报错信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回dto
	Data *SupplierBasisInfoDto `json:"data,omitempty" xml:"data,omitempty"`
	// 报错code
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// ts
	Ts int64 `json:"ts,omitempty" xml:"ts,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolJsonResponse = sync.Pool{
	New: func() any {
		return new(JsonResponse)
	},
}

// GetJsonResponse() 从对象池中获取JsonResponse
func GetJsonResponse() *JsonResponse {
	return poolJsonResponse.Get().(*JsonResponse)
}

// ReleaseJsonResponse 释放JsonResponse
func ReleaseJsonResponse(v *JsonResponse) {
	v.ErrMsg = ""
	v.Data = nil
	v.ErrCode = 0
	v.Ts = 0
	v.Success = false
	poolJsonResponse.Put(v)
}
