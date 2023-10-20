package ascp

import (
	"sync"
)

// MaterialChangeResponse 结构体
type MaterialChangeResponse struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 处理结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMaterialChangeResponse = sync.Pool{
	New: func() any {
		return new(MaterialChangeResponse)
	},
}

// GetMaterialChangeResponse() 从对象池中获取MaterialChangeResponse
func GetMaterialChangeResponse() *MaterialChangeResponse {
	return poolMaterialChangeResponse.Get().(*MaterialChangeResponse)
}

// ReleaseMaterialChangeResponse 释放MaterialChangeResponse
func ReleaseMaterialChangeResponse(v *MaterialChangeResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolMaterialChangeResponse.Put(v)
}
