package examination

import (
	"sync"
)

// ReserveStatusResultResponse 结构体
type ReserveStatusResultResponse struct {
	// 业务响应code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 正文
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}

var poolReserveStatusResultResponse = sync.Pool{
	New: func() any {
		return new(ReserveStatusResultResponse)
	},
}

// GetReserveStatusResultResponse() 从对象池中获取ReserveStatusResultResponse
func GetReserveStatusResultResponse() *ReserveStatusResultResponse {
	return poolReserveStatusResultResponse.Get().(*ReserveStatusResultResponse)
}

// ReleaseReserveStatusResultResponse 释放ReserveStatusResultResponse
func ReleaseReserveStatusResultResponse(v *ReserveStatusResultResponse) {
	v.ResponseCode = ""
	v.Msg = ""
	poolReserveStatusResultResponse.Put(v)
}
