package ascp

import (
	"sync"
)

// LspTopResponse 结构体
type LspTopResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 上传结果
	Data *MediaResourceDto `json:"data,omitempty" xml:"data,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolLspTopResponse = sync.Pool{
	New: func() any {
		return new(LspTopResponse)
	},
}

// GetLspTopResponse() 从对象池中获取LspTopResponse
func GetLspTopResponse() *LspTopResponse {
	return poolLspTopResponse.Get().(*LspTopResponse)
}

// ReleaseLspTopResponse 释放LspTopResponse
func ReleaseLspTopResponse(v *LspTopResponse) {
	v.Code = ""
	v.Message = ""
	v.Result = ""
	v.TraceId = ""
	v.Data = nil
	v.Success = false
	poolLspTopResponse.Put(v)
}
