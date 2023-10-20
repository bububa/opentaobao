package examination

import (
	"sync"
)

// ModifyResultResponse 结构体
type ModifyResultResponse struct {
	// 结果
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 结果
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}

var poolModifyResultResponse = sync.Pool{
	New: func() any {
		return new(ModifyResultResponse)
	},
}

// GetModifyResultResponse() 从对象池中获取ModifyResultResponse
func GetModifyResultResponse() *ModifyResultResponse {
	return poolModifyResultResponse.Get().(*ModifyResultResponse)
}

// ReleaseModifyResultResponse 释放ModifyResultResponse
func ReleaseModifyResultResponse(v *ModifyResultResponse) {
	v.ResponseCode = ""
	v.Msg = ""
	poolModifyResultResponse.Put(v)
}
