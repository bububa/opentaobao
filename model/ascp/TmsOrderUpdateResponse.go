package ascp

import (
	"sync"
)

// TmsOrderUpdateResponse 结构体
type TmsOrderUpdateResponse struct {
	// 业务错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否为业务异常在流程中心场景下使用
	BizException bool `json:"biz_exception,omitempty" xml:"biz_exception,omitempty"`
	// 业务调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmsOrderUpdateResponse = sync.Pool{
	New: func() any {
		return new(TmsOrderUpdateResponse)
	},
}

// GetTmsOrderUpdateResponse() 从对象池中获取TmsOrderUpdateResponse
func GetTmsOrderUpdateResponse() *TmsOrderUpdateResponse {
	return poolTmsOrderUpdateResponse.Get().(*TmsOrderUpdateResponse)
}

// ReleaseTmsOrderUpdateResponse 释放TmsOrderUpdateResponse
func ReleaseTmsOrderUpdateResponse(v *TmsOrderUpdateResponse) {
	v.Code = ""
	v.Message = ""
	v.BizException = false
	v.Success = false
	poolTmsOrderUpdateResponse.Put(v)
}
