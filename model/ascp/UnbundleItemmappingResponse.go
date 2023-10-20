package ascp

import (
	"sync"
)

// UnbundleItemmappingResponse 结构体
type UnbundleItemmappingResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务处理结果
	Data *UnbundleItemmappingResult `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolUnbundleItemmappingResponse = sync.Pool{
	New: func() any {
		return new(UnbundleItemmappingResponse)
	},
}

// GetUnbundleItemmappingResponse() 从对象池中获取UnbundleItemmappingResponse
func GetUnbundleItemmappingResponse() *UnbundleItemmappingResponse {
	return poolUnbundleItemmappingResponse.Get().(*UnbundleItemmappingResponse)
}

// ReleaseUnbundleItemmappingResponse 释放UnbundleItemmappingResponse
func ReleaseUnbundleItemmappingResponse(v *UnbundleItemmappingResponse) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolUnbundleItemmappingResponse.Put(v)
}
