package qimen

import (
	"sync"
)

// TaobaoQimenShopSynchronizeRequest 结构体
type TaobaoQimenShopSynchronizeRequest struct {
	// add,update, 必填
	ActionType string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 店铺
	Shop *Shop `json:"shop,omitempty" xml:"shop,omitempty"`
}

var poolTaobaoQimenShopSynchronizeRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenShopSynchronizeRequest)
	},
}

// GetTaobaoQimenShopSynchronizeRequest() 从对象池中获取TaobaoQimenShopSynchronizeRequest
func GetTaobaoQimenShopSynchronizeRequest() *TaobaoQimenShopSynchronizeRequest {
	return poolTaobaoQimenShopSynchronizeRequest.Get().(*TaobaoQimenShopSynchronizeRequest)
}

// ReleaseTaobaoQimenShopSynchronizeRequest 释放TaobaoQimenShopSynchronizeRequest
func ReleaseTaobaoQimenShopSynchronizeRequest(v *TaobaoQimenShopSynchronizeRequest) {
	v.ActionType = ""
	v.Shop = nil
	poolTaobaoQimenShopSynchronizeRequest.Put(v)
}
