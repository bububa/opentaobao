package alicom

import (
	"sync"
)

// AlibabaAliqinAxbVendorSmsInterceptResponse 结构体
type AlibabaAliqinAxbVendorSmsInterceptResponse struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应的业务CODE：OK代表请求成功，非OK的错误码建议进行重试
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 此字段忽略，只用判断code是否为OK
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

var poolAlibabaAliqinAxbVendorSmsInterceptResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinAxbVendorSmsInterceptResponse)
	},
}

// GetAlibabaAliqinAxbVendorSmsInterceptResponse() 从对象池中获取AlibabaAliqinAxbVendorSmsInterceptResponse
func GetAlibabaAliqinAxbVendorSmsInterceptResponse() *AlibabaAliqinAxbVendorSmsInterceptResponse {
	return poolAlibabaAliqinAxbVendorSmsInterceptResponse.Get().(*AlibabaAliqinAxbVendorSmsInterceptResponse)
}

// ReleaseAlibabaAliqinAxbVendorSmsInterceptResponse 释放AlibabaAliqinAxbVendorSmsInterceptResponse
func ReleaseAlibabaAliqinAxbVendorSmsInterceptResponse(v *AlibabaAliqinAxbVendorSmsInterceptResponse) {
	v.Message = ""
	v.Code = ""
	v.Module = false
	poolAlibabaAliqinAxbVendorSmsInterceptResponse.Put(v)
}
