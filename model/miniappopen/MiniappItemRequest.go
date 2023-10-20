package miniappopen

import (
	"sync"
)

// MiniappItemRequest 结构体
type MiniappItemRequest struct {
	// 商品id列表
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 小程序小游戏appId
	AppId int64 `json:"app_id,omitempty" xml:"app_id,omitempty"`
}

var poolMiniappItemRequest = sync.Pool{
	New: func() any {
		return new(MiniappItemRequest)
	},
}

// GetMiniappItemRequest() 从对象池中获取MiniappItemRequest
func GetMiniappItemRequest() *MiniappItemRequest {
	return poolMiniappItemRequest.Get().(*MiniappItemRequest)
}

// ReleaseMiniappItemRequest 释放MiniappItemRequest
func ReleaseMiniappItemRequest(v *MiniappItemRequest) {
	v.ItemIds = v.ItemIds[:0]
	v.AppId = 0
	poolMiniappItemRequest.Put(v)
}
