package moscm

import (
	"sync"
)

// IsvOutboundRequestItemDto 结构体
type IsvOutboundRequestItemDto struct {
	// 外部id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 数量
	Quantity float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolIsvOutboundRequestItemDto = sync.Pool{
	New: func() any {
		return new(IsvOutboundRequestItemDto)
	},
}

// GetIsvOutboundRequestItemDto() 从对象池中获取IsvOutboundRequestItemDto
func GetIsvOutboundRequestItemDto() *IsvOutboundRequestItemDto {
	return poolIsvOutboundRequestItemDto.Get().(*IsvOutboundRequestItemDto)
}

// ReleaseIsvOutboundRequestItemDto 释放IsvOutboundRequestItemDto
func ReleaseIsvOutboundRequestItemDto(v *IsvOutboundRequestItemDto) {
	v.OutId = ""
	v.Quantity = 0
	poolIsvOutboundRequestItemDto.Put(v)
}
