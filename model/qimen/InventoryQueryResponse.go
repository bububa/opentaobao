package qimen

import (
	"sync"
)

// InventoryQueryResponse 结构体
type InventoryQueryResponse struct {
	// 商品的库存信息列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolInventoryQueryResponse = sync.Pool{
	New: func() any {
		return new(InventoryQueryResponse)
	},
}

// GetInventoryQueryResponse() 从对象池中获取InventoryQueryResponse
func GetInventoryQueryResponse() *InventoryQueryResponse {
	return poolInventoryQueryResponse.Get().(*InventoryQueryResponse)
}

// ReleaseInventoryQueryResponse 释放InventoryQueryResponse
func ReleaseInventoryQueryResponse(v *InventoryQueryResponse) {
	v.Items = v.Items[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolInventoryQueryResponse.Put(v)
}
