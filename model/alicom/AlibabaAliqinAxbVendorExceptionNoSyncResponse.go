package alicom

import (
	"sync"
)

// AlibabaAliqinAxbVendorExceptionNoSyncResponse 结构体
type AlibabaAliqinAxbVendorExceptionNoSyncResponse struct {
	// 错误信息,OK代表受理成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码,OK代表受理成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// module
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

var poolAlibabaAliqinAxbVendorExceptionNoSyncResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinAxbVendorExceptionNoSyncResponse)
	},
}

// GetAlibabaAliqinAxbVendorExceptionNoSyncResponse() 从对象池中获取AlibabaAliqinAxbVendorExceptionNoSyncResponse
func GetAlibabaAliqinAxbVendorExceptionNoSyncResponse() *AlibabaAliqinAxbVendorExceptionNoSyncResponse {
	return poolAlibabaAliqinAxbVendorExceptionNoSyncResponse.Get().(*AlibabaAliqinAxbVendorExceptionNoSyncResponse)
}

// ReleaseAlibabaAliqinAxbVendorExceptionNoSyncResponse 释放AlibabaAliqinAxbVendorExceptionNoSyncResponse
func ReleaseAlibabaAliqinAxbVendorExceptionNoSyncResponse(v *AlibabaAliqinAxbVendorExceptionNoSyncResponse) {
	v.Message = ""
	v.Code = ""
	v.Module = false
	poolAlibabaAliqinAxbVendorExceptionNoSyncResponse.Put(v)
}
