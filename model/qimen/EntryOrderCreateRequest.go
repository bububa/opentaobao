package qimen

import (
	"sync"
)

// EntryOrderCreateRequest 结构体
type EntryOrderCreateRequest struct {
	// 入库单详情
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 入库单信息
	EntryOrder *EntryOrder `json:"entryOrder,omitempty" xml:"entryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenEntryorderCreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolEntryOrderCreateRequest = sync.Pool{
	New: func() any {
		return new(EntryOrderCreateRequest)
	},
}

// GetEntryOrderCreateRequest() 从对象池中获取EntryOrderCreateRequest
func GetEntryOrderCreateRequest() *EntryOrderCreateRequest {
	return poolEntryOrderCreateRequest.Get().(*EntryOrderCreateRequest)
}

// ReleaseEntryOrderCreateRequest 释放EntryOrderCreateRequest
func ReleaseEntryOrderCreateRequest(v *EntryOrderCreateRequest) {
	v.OrderLines = v.OrderLines[:0]
	v.EntryOrder = nil
	v.ExtendProps = nil
	poolEntryOrderCreateRequest.Put(v)
}
