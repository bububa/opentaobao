package moscm

import (
	"sync"
)

// IsvStockAdjustRequestDto 结构体
type IsvStockAdjustRequestDto struct {
	// 入库项（最大列表长度：20）
	InboundItems []IsvInboundRequestItemDto `json:"inbound_items,omitempty" xml:"inbound_items>isv_inbound_request_item_dto,omitempty"`
	// 出库项（最大列表长度：20）
	OutboundItems []IsvOutboundRequestItemDto `json:"outbound_items,omitempty" xml:"outbound_items>isv_outbound_request_item_dto,omitempty"`
	// 专柜Id
	CounterId string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
	// 外部专柜Id（两个专柜id字段至少填写一个，如果同时填写，已counter_id为准）
	OutCounterId string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
	// 外部单号
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 备注
	Remarks string `json:"remarks,omitempty" xml:"remarks,omitempty"`
}

var poolIsvStockAdjustRequestDto = sync.Pool{
	New: func() any {
		return new(IsvStockAdjustRequestDto)
	},
}

// GetIsvStockAdjustRequestDto() 从对象池中获取IsvStockAdjustRequestDto
func GetIsvStockAdjustRequestDto() *IsvStockAdjustRequestDto {
	return poolIsvStockAdjustRequestDto.Get().(*IsvStockAdjustRequestDto)
}

// ReleaseIsvStockAdjustRequestDto 释放IsvStockAdjustRequestDto
func ReleaseIsvStockAdjustRequestDto(v *IsvStockAdjustRequestDto) {
	v.InboundItems = v.InboundItems[:0]
	v.OutboundItems = v.OutboundItems[:0]
	v.CounterId = ""
	v.OutCounterId = ""
	v.OutId = ""
	v.Remarks = ""
	poolIsvStockAdjustRequestDto.Put(v)
}
