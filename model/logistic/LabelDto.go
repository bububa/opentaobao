package logistic

import (
	"sync"
)

// LabelDto 结构体
type LabelDto struct {
	// 物流订单号
	LogisticsOrderId string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
	// 面单链接
	LabelUrl string `json:"label_url,omitempty" xml:"label_url,omitempty"`
}

var poolLabelDto = sync.Pool{
	New: func() any {
		return new(LabelDto)
	},
}

// GetLabelDto() 从对象池中获取LabelDto
func GetLabelDto() *LabelDto {
	return poolLabelDto.Get().(*LabelDto)
}

// ReleaseLabelDto 释放LabelDto
func ReleaseLabelDto(v *LabelDto) {
	v.LogisticsOrderId = ""
	v.LabelUrl = ""
	poolLabelDto.Put(v)
}
