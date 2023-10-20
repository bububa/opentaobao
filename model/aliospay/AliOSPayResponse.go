package aliospay

import (
	"sync"
)

// AliOSPayResponse 结构体
type AliOSPayResponse struct {
	// 请求唯一id，不可重复，服务端会根据此参数防重放
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 业务数据
	Data *PeriodAgreementPayResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAliOSPayResponse = sync.Pool{
	New: func() any {
		return new(AliOSPayResponse)
	},
}

// GetAliOSPayResponse() 从对象池中获取AliOSPayResponse
func GetAliOSPayResponse() *AliOSPayResponse {
	return poolAliOSPayResponse.Get().(*AliOSPayResponse)
}

// ReleaseAliOSPayResponse 释放AliOSPayResponse
func ReleaseAliOSPayResponse(v *AliOSPayResponse) {
	v.TraceId = ""
	v.Message = ""
	v.Code = 0
	v.Data = nil
	poolAliOSPayResponse.Put(v)
}
