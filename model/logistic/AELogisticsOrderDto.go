package logistic

import (
	"sync"
)

// AELogisticsOrderDto 结构体
type AELogisticsOrderDto struct {
	// Shipment order id created. AE will save relationship with logistics provider&#39;s shipment order_id
	LogisticsOrderId string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
}

var poolAELogisticsOrderDto = sync.Pool{
	New: func() any {
		return new(AELogisticsOrderDto)
	},
}

// GetAELogisticsOrderDto() 从对象池中获取AELogisticsOrderDto
func GetAELogisticsOrderDto() *AELogisticsOrderDto {
	return poolAELogisticsOrderDto.Get().(*AELogisticsOrderDto)
}

// ReleaseAELogisticsOrderDto 释放AELogisticsOrderDto
func ReleaseAELogisticsOrderDto(v *AELogisticsOrderDto) {
	v.LogisticsOrderId = ""
	poolAELogisticsOrderDto.Put(v)
}
