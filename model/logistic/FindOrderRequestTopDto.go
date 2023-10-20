package logistic

import (
	"sync"
)

// FindOrderRequestTopDto 结构体
type FindOrderRequestTopDto struct {
	// Shipment order id created. AE will save relationship with logistics provider&#39;s shipment order_id
	LogisticsOrderId string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
}

var poolFindOrderRequestTopDto = sync.Pool{
	New: func() any {
		return new(FindOrderRequestTopDto)
	},
}

// GetFindOrderRequestTopDto() 从对象池中获取FindOrderRequestTopDto
func GetFindOrderRequestTopDto() *FindOrderRequestTopDto {
	return poolFindOrderRequestTopDto.Get().(*FindOrderRequestTopDto)
}

// ReleaseFindOrderRequestTopDto 释放FindOrderRequestTopDto
func ReleaseFindOrderRequestTopDto(v *FindOrderRequestTopDto) {
	v.LogisticsOrderId = ""
	poolFindOrderRequestTopDto.Put(v)
}
