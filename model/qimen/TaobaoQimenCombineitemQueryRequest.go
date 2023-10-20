package qimen

import (
	"sync"
)

// TaobaoQimenCombineitemQueryRequest 结构体
type TaobaoQimenCombineitemQueryRequest struct {
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenCombineitemQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolTaobaoQimenCombineitemQueryRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenCombineitemQueryRequest)
	},
}

// GetTaobaoQimenCombineitemQueryRequest() 从对象池中获取TaobaoQimenCombineitemQueryRequest
func GetTaobaoQimenCombineitemQueryRequest() *TaobaoQimenCombineitemQueryRequest {
	return poolTaobaoQimenCombineitemQueryRequest.Get().(*TaobaoQimenCombineitemQueryRequest)
}

// ReleaseTaobaoQimenCombineitemQueryRequest 释放TaobaoQimenCombineitemQueryRequest
func ReleaseTaobaoQimenCombineitemQueryRequest(v *TaobaoQimenCombineitemQueryRequest) {
	v.OwnerCode = ""
	v.ItemId = ""
	v.ExtendProps = nil
	poolTaobaoQimenCombineitemQueryRequest.Put(v)
}
