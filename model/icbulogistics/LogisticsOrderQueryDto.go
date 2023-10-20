package icbulogistics

import (
	"sync"
)

// LogisticsOrderQueryDto 结构体
type LogisticsOrderQueryDto struct {
	// 物流单号
	OrderNumber string `json:"order_number,omitempty" xml:"order_number,omitempty"`
}

var poolLogisticsOrderQueryDto = sync.Pool{
	New: func() any {
		return new(LogisticsOrderQueryDto)
	},
}

// GetLogisticsOrderQueryDto() 从对象池中获取LogisticsOrderQueryDto
func GetLogisticsOrderQueryDto() *LogisticsOrderQueryDto {
	return poolLogisticsOrderQueryDto.Get().(*LogisticsOrderQueryDto)
}

// ReleaseLogisticsOrderQueryDto 释放LogisticsOrderQueryDto
func ReleaseLogisticsOrderQueryDto(v *LogisticsOrderQueryDto) {
	v.OrderNumber = ""
	poolLogisticsOrderQueryDto.Put(v)
}
