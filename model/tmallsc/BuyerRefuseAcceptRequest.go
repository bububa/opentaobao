package tmallsc

import (
	"sync"
)

// BuyerRefuseAcceptRequest 结构体
type BuyerRefuseAcceptRequest struct {
	// 拒收备注
	RefuseMemo string `json:"refuse_memo,omitempty" xml:"refuse_memo,omitempty"`
	// 工单ID
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}

var poolBuyerRefuseAcceptRequest = sync.Pool{
	New: func() any {
		return new(BuyerRefuseAcceptRequest)
	},
}

// GetBuyerRefuseAcceptRequest() 从对象池中获取BuyerRefuseAcceptRequest
func GetBuyerRefuseAcceptRequest() *BuyerRefuseAcceptRequest {
	return poolBuyerRefuseAcceptRequest.Get().(*BuyerRefuseAcceptRequest)
}

// ReleaseBuyerRefuseAcceptRequest 释放BuyerRefuseAcceptRequest
func ReleaseBuyerRefuseAcceptRequest(v *BuyerRefuseAcceptRequest) {
	v.RefuseMemo = ""
	v.WorkcardId = 0
	poolBuyerRefuseAcceptRequest.Put(v)
}
