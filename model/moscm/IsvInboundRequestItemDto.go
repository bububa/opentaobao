package moscm

import (
	"sync"
)

// IsvInboundRequestItemDto 结构体
type IsvInboundRequestItemDto struct {
	// 库位编号
	LocationId string `json:"location_id,omitempty" xml:"location_id,omitempty"`
	// 外部id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 获取或设置货物数量计量单位
	Quantity float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 单价
	UnitPrice float64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}

var poolIsvInboundRequestItemDto = sync.Pool{
	New: func() any {
		return new(IsvInboundRequestItemDto)
	},
}

// GetIsvInboundRequestItemDto() 从对象池中获取IsvInboundRequestItemDto
func GetIsvInboundRequestItemDto() *IsvInboundRequestItemDto {
	return poolIsvInboundRequestItemDto.Get().(*IsvInboundRequestItemDto)
}

// ReleaseIsvInboundRequestItemDto 释放IsvInboundRequestItemDto
func ReleaseIsvInboundRequestItemDto(v *IsvInboundRequestItemDto) {
	v.LocationId = ""
	v.OutId = ""
	v.Unit = ""
	v.Quantity = 0
	v.UnitPrice = 0
	poolIsvInboundRequestItemDto.Put(v)
}
