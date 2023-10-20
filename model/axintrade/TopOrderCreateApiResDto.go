package axintrade

import (
	"sync"
)

// TopOrderCreateApiResDto 结构体
type TopOrderCreateApiResDto struct {
	// 采购单号
	PurchaseSubOrderId int64 `json:"purchase_sub_order_id,omitempty" xml:"purchase_sub_order_id,omitempty"`
}

var poolTopOrderCreateApiResDto = sync.Pool{
	New: func() any {
		return new(TopOrderCreateApiResDto)
	},
}

// GetTopOrderCreateApiResDto() 从对象池中获取TopOrderCreateApiResDto
func GetTopOrderCreateApiResDto() *TopOrderCreateApiResDto {
	return poolTopOrderCreateApiResDto.Get().(*TopOrderCreateApiResDto)
}

// ReleaseTopOrderCreateApiResDto 释放TopOrderCreateApiResDto
func ReleaseTopOrderCreateApiResDto(v *TopOrderCreateApiResDto) {
	v.PurchaseSubOrderId = 0
	poolTopOrderCreateApiResDto.Put(v)
}
