package tblogistics

import (
	"sync"
)

// ConsignDto 结构体
type ConsignDto struct {
	// 发货文案描述
	ConsignDesc string `json:"consign_desc,omitempty" xml:"consign_desc,omitempty"`
	// 物流发货单号
	LpOrderId int64 `json:"lp_order_id,omitempty" xml:"lp_order_id,omitempty"`
}

var poolConsignDto = sync.Pool{
	New: func() any {
		return new(ConsignDto)
	},
}

// GetConsignDto() 从对象池中获取ConsignDto
func GetConsignDto() *ConsignDto {
	return poolConsignDto.Get().(*ConsignDto)
}

// ReleaseConsignDto 释放ConsignDto
func ReleaseConsignDto(v *ConsignDto) {
	v.ConsignDesc = ""
	v.LpOrderId = 0
	poolConsignDto.Put(v)
}
