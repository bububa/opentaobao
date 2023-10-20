package alicom

import (
	"sync"
)

// AlibabaAliqinAxbVendorCallControlResponse 结构体
type AlibabaAliqinAxbVendorCallControlResponse struct {
	// 转呼控制msg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// module
	ControlRespDto *ControlRespDto `json:"control_resp_dto,omitempty" xml:"control_resp_dto,omitempty"`
}

var poolAlibabaAliqinAxbVendorCallControlResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinAxbVendorCallControlResponse)
	},
}

// GetAlibabaAliqinAxbVendorCallControlResponse() 从对象池中获取AlibabaAliqinAxbVendorCallControlResponse
func GetAlibabaAliqinAxbVendorCallControlResponse() *AlibabaAliqinAxbVendorCallControlResponse {
	return poolAlibabaAliqinAxbVendorCallControlResponse.Get().(*AlibabaAliqinAxbVendorCallControlResponse)
}

// ReleaseAlibabaAliqinAxbVendorCallControlResponse 释放AlibabaAliqinAxbVendorCallControlResponse
func ReleaseAlibabaAliqinAxbVendorCallControlResponse(v *AlibabaAliqinAxbVendorCallControlResponse) {
	v.Message = ""
	v.Code = ""
	v.ControlRespDto = nil
	poolAlibabaAliqinAxbVendorCallControlResponse.Put(v)
}
