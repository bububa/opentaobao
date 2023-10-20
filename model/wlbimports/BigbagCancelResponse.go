package wlbimports

import (
	"sync"
)

// BigbagCancelResponse 结构体
type BigbagCancelResponse struct {
	// 取消是否成功
	CancelFlag bool `json:"cancel_flag,omitempty" xml:"cancel_flag,omitempty"`
}

var poolBigbagCancelResponse = sync.Pool{
	New: func() any {
		return new(BigbagCancelResponse)
	},
}

// GetBigbagCancelResponse() 从对象池中获取BigbagCancelResponse
func GetBigbagCancelResponse() *BigbagCancelResponse {
	return poolBigbagCancelResponse.Get().(*BigbagCancelResponse)
}

// ReleaseBigbagCancelResponse 释放BigbagCancelResponse
func ReleaseBigbagCancelResponse(v *BigbagCancelResponse) {
	v.CancelFlag = false
	poolBigbagCancelResponse.Put(v)
}
