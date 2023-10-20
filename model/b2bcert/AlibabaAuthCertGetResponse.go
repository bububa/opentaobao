package b2bcert

import (
	"sync"
)

// AlibabaAuthCertGetResponse 结构体
type AlibabaAuthCertGetResponse struct {
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAuthCertGetResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAuthCertGetResponse)
	},
}

// GetAlibabaAuthCertGetResponse() 从对象池中获取AlibabaAuthCertGetResponse
func GetAlibabaAuthCertGetResponse() *AlibabaAuthCertGetResponse {
	return poolAlibabaAuthCertGetResponse.Get().(*AlibabaAuthCertGetResponse)
}

// ReleaseAlibabaAuthCertGetResponse 释放AlibabaAuthCertGetResponse
func ReleaseAlibabaAuthCertGetResponse(v *AlibabaAuthCertGetResponse) {
	v.Data = ""
	v.Code = ""
	v.Msg = ""
	v.Success = false
	poolAlibabaAuthCertGetResponse.Put(v)
}
