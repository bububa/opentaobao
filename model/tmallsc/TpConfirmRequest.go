package tmallsc

import (
	"sync"
)

// TpConfirmRequest 结构体
type TpConfirmRequest struct {
	// 工单id
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}

var poolTpConfirmRequest = sync.Pool{
	New: func() any {
		return new(TpConfirmRequest)
	},
}

// GetTpConfirmRequest() 从对象池中获取TpConfirmRequest
func GetTpConfirmRequest() *TpConfirmRequest {
	return poolTpConfirmRequest.Get().(*TpConfirmRequest)
}

// ReleaseTpConfirmRequest 释放TpConfirmRequest
func ReleaseTpConfirmRequest(v *TpConfirmRequest) {
	v.WorkcardId = 0
	poolTpConfirmRequest.Put(v)
}
