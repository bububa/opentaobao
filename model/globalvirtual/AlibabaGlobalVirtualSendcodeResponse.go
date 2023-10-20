package globalvirtual

import (
	"sync"
)

// AlibabaGlobalVirtualSendcodeResponse 结构体
type AlibabaGlobalVirtualSendcodeResponse struct {
	// error code
	ErrorCode *ErrorCode `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// request result
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// send code result
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
	// request is repeated
	Repeated bool `json:"repeated,omitempty" xml:"repeated,omitempty"`
	// request need retry
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

var poolAlibabaGlobalVirtualSendcodeResponse = sync.Pool{
	New: func() any {
		return new(AlibabaGlobalVirtualSendcodeResponse)
	},
}

// GetAlibabaGlobalVirtualSendcodeResponse() 从对象池中获取AlibabaGlobalVirtualSendcodeResponse
func GetAlibabaGlobalVirtualSendcodeResponse() *AlibabaGlobalVirtualSendcodeResponse {
	return poolAlibabaGlobalVirtualSendcodeResponse.Get().(*AlibabaGlobalVirtualSendcodeResponse)
}

// ReleaseAlibabaGlobalVirtualSendcodeResponse 释放AlibabaGlobalVirtualSendcodeResponse
func ReleaseAlibabaGlobalVirtualSendcodeResponse(v *AlibabaGlobalVirtualSendcodeResponse) {
	v.ErrorCode = nil
	v.Success = false
	v.Module = false
	v.Repeated = false
	v.Retry = false
	poolAlibabaGlobalVirtualSendcodeResponse.Put(v)
}
