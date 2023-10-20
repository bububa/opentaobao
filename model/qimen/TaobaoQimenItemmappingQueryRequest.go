package qimen

import (
	"sync"
)

// TaobaoQimenItemmappingQueryRequest 结构体
type TaobaoQimenItemmappingQueryRequest struct {
	// 奇门仓储字段,C123,string(50),,
	QueryType string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	// 奇门仓储字段,C123,string(50),,
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段,C123,string(50),,
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 奇门仓储字段,C123,string(50),,
	ShopItemId string `json:"shopItemId,omitempty" xml:"shopItemId,omitempty"`
	// 奇门仓储字段,C123,string(50),,
	SkuId string `json:"skuId,omitempty" xml:"skuId,omitempty"`
}

var poolTaobaoQimenItemmappingQueryRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemmappingQueryRequest)
	},
}

// GetTaobaoQimenItemmappingQueryRequest() 从对象池中获取TaobaoQimenItemmappingQueryRequest
func GetTaobaoQimenItemmappingQueryRequest() *TaobaoQimenItemmappingQueryRequest {
	return poolTaobaoQimenItemmappingQueryRequest.Get().(*TaobaoQimenItemmappingQueryRequest)
}

// ReleaseTaobaoQimenItemmappingQueryRequest 释放TaobaoQimenItemmappingQueryRequest
func ReleaseTaobaoQimenItemmappingQueryRequest(v *TaobaoQimenItemmappingQueryRequest) {
	v.QueryType = ""
	v.OwnerCode = ""
	v.ItemId = ""
	v.ShopItemId = ""
	v.SkuId = ""
	poolTaobaoQimenItemmappingQueryRequest.Put(v)
}
