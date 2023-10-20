package fundplatform

import (
	"sync"
)

// CardMakingInfoQueryRequest 结构体
type CardMakingInfoQueryRequest struct {
	// 卡号列表
	CardNos []string `json:"card_nos,omitempty" xml:"card_nos>string,omitempty"`
	// 已废弃,可以填写固定值test
	Signture string `json:"signture,omitempty" xml:"signture,omitempty"`
	// 页大小,最大500
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页，从1开始
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 子制卡单ID
	CardOrderId int64 `json:"card_order_id,omitempty" xml:"card_order_id,omitempty"`
}

var poolCardMakingInfoQueryRequest = sync.Pool{
	New: func() any {
		return new(CardMakingInfoQueryRequest)
	},
}

// GetCardMakingInfoQueryRequest() 从对象池中获取CardMakingInfoQueryRequest
func GetCardMakingInfoQueryRequest() *CardMakingInfoQueryRequest {
	return poolCardMakingInfoQueryRequest.Get().(*CardMakingInfoQueryRequest)
}

// ReleaseCardMakingInfoQueryRequest 释放CardMakingInfoQueryRequest
func ReleaseCardMakingInfoQueryRequest(v *CardMakingInfoQueryRequest) {
	v.CardNos = v.CardNos[:0]
	v.Signture = ""
	v.PageSize = 0
	v.CurrentPage = 0
	v.CardOrderId = 0
	poolCardMakingInfoQueryRequest.Put(v)
}
