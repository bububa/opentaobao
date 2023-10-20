package alicom

import (
	"sync"
)

// AlibabaAliqinAxbVendorHeartBeatResponse 结构体
type AlibabaAliqinAxbVendorHeartBeatResponse struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// module
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

var poolAlibabaAliqinAxbVendorHeartBeatResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinAxbVendorHeartBeatResponse)
	},
}

// GetAlibabaAliqinAxbVendorHeartBeatResponse() 从对象池中获取AlibabaAliqinAxbVendorHeartBeatResponse
func GetAlibabaAliqinAxbVendorHeartBeatResponse() *AlibabaAliqinAxbVendorHeartBeatResponse {
	return poolAlibabaAliqinAxbVendorHeartBeatResponse.Get().(*AlibabaAliqinAxbVendorHeartBeatResponse)
}

// ReleaseAlibabaAliqinAxbVendorHeartBeatResponse 释放AlibabaAliqinAxbVendorHeartBeatResponse
func ReleaseAlibabaAliqinAxbVendorHeartBeatResponse(v *AlibabaAliqinAxbVendorHeartBeatResponse) {
	v.Message = ""
	v.Code = ""
	v.Module = false
	poolAlibabaAliqinAxbVendorHeartBeatResponse.Put(v)
}
