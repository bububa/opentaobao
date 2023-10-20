package ascp

import (
	"sync"
)

// TmsOrderConfirmResponse 结构体
type TmsOrderConfirmResponse struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 业务请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否业务异常
	BizException bool `json:"biz_exception,omitempty" xml:"biz_exception,omitempty"`
}

var poolTmsOrderConfirmResponse = sync.Pool{
	New: func() any {
		return new(TmsOrderConfirmResponse)
	},
}

// GetTmsOrderConfirmResponse() 从对象池中获取TmsOrderConfirmResponse
func GetTmsOrderConfirmResponse() *TmsOrderConfirmResponse {
	return poolTmsOrderConfirmResponse.Get().(*TmsOrderConfirmResponse)
}

// ReleaseTmsOrderConfirmResponse 释放TmsOrderConfirmResponse
func ReleaseTmsOrderConfirmResponse(v *TmsOrderConfirmResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.BizException = false
	poolTmsOrderConfirmResponse.Put(v)
}
