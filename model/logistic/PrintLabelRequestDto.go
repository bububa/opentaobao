package logistic

import (
	"sync"
)

// PrintLabelRequestDto 结构体
type PrintLabelRequestDto struct {
	// 物流订单号
	LogisticsOrderId string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
}

var poolPrintLabelRequestDto = sync.Pool{
	New: func() any {
		return new(PrintLabelRequestDto)
	},
}

// GetPrintLabelRequestDto() 从对象池中获取PrintLabelRequestDto
func GetPrintLabelRequestDto() *PrintLabelRequestDto {
	return poolPrintLabelRequestDto.Get().(*PrintLabelRequestDto)
}

// ReleasePrintLabelRequestDto 释放PrintLabelRequestDto
func ReleasePrintLabelRequestDto(v *PrintLabelRequestDto) {
	v.LogisticsOrderId = ""
	poolPrintLabelRequestDto.Put(v)
}
