package qimen

import (
	"sync"
)

// TaobaoQimenInventoryreserveCancelResponse 结构体
type TaobaoQimenInventoryreserveCancelResponse struct {
	// 奇门仓储字段
	ItemInventories []ItemInventory `json:"itemInventories,omitempty" xml:"itemInventories>item_inventory,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 奇门仓储字段
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 奇门仓储字段
	IsRetry string `json:"isRetry,omitempty" xml:"isRetry,omitempty"`
}

var poolTaobaoQimenInventoryreserveCancelResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventoryreserveCancelResponse)
	},
}

// GetTaobaoQimenInventoryreserveCancelResponse() 从对象池中获取TaobaoQimenInventoryreserveCancelResponse
func GetTaobaoQimenInventoryreserveCancelResponse() *TaobaoQimenInventoryreserveCancelResponse {
	return poolTaobaoQimenInventoryreserveCancelResponse.Get().(*TaobaoQimenInventoryreserveCancelResponse)
}

// ReleaseTaobaoQimenInventoryreserveCancelResponse 释放TaobaoQimenInventoryreserveCancelResponse
func ReleaseTaobaoQimenInventoryreserveCancelResponse(v *TaobaoQimenInventoryreserveCancelResponse) {
	v.ItemInventories = v.ItemInventories[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.OrderCode = ""
	v.IsRetry = ""
	poolTaobaoQimenInventoryreserveCancelResponse.Put(v)
}
