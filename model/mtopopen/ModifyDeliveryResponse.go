package mtopopen

import (
	"sync"
)

// ModifyDeliveryResponse 结构体
type ModifyDeliveryResponse struct {
	// 信息回传错误码(成功情况无需关注)
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 信息回传错误原因(成功情况无需关注)
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 信息回传是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolModifyDeliveryResponse = sync.Pool{
	New: func() any {
		return new(ModifyDeliveryResponse)
	},
}

// GetModifyDeliveryResponse() 从对象池中获取ModifyDeliveryResponse
func GetModifyDeliveryResponse() *ModifyDeliveryResponse {
	return poolModifyDeliveryResponse.Get().(*ModifyDeliveryResponse)
}

// ReleaseModifyDeliveryResponse 释放ModifyDeliveryResponse
func ReleaseModifyDeliveryResponse(v *ModifyDeliveryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolModifyDeliveryResponse.Put(v)
}
