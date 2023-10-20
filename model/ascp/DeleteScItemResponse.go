package ascp

import (
	"sync"
)

// DeleteScItemResponse 结构体
type DeleteScItemResponse struct {
	// 调用链路ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回结果
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDeleteScItemResponse = sync.Pool{
	New: func() any {
		return new(DeleteScItemResponse)
	},
}

// GetDeleteScItemResponse() 从对象池中获取DeleteScItemResponse
func GetDeleteScItemResponse() *DeleteScItemResponse {
	return poolDeleteScItemResponse.Get().(*DeleteScItemResponse)
}

// ReleaseDeleteScItemResponse 释放DeleteScItemResponse
func ReleaseDeleteScItemResponse(v *DeleteScItemResponse) {
	v.TraceId = ""
	v.Code = ""
	v.Data = ""
	v.Message = ""
	v.Success = false
	poolDeleteScItemResponse.Put(v)
}
