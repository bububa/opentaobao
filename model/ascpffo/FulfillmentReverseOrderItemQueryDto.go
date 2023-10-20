package ascpffo

import (
	"sync"
)

// FulfillmentReverseOrderItemQueryDto 结构体
type FulfillmentReverseOrderItemQueryDto struct {
	// 履约单号
	FulfillmentOrderNo string `json:"fulfillment_order_no,omitempty" xml:"fulfillment_order_no,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

var poolFulfillmentReverseOrderItemQueryDto = sync.Pool{
	New: func() any {
		return new(FulfillmentReverseOrderItemQueryDto)
	},
}

// GetFulfillmentReverseOrderItemQueryDto() 从对象池中获取FulfillmentReverseOrderItemQueryDto
func GetFulfillmentReverseOrderItemQueryDto() *FulfillmentReverseOrderItemQueryDto {
	return poolFulfillmentReverseOrderItemQueryDto.Get().(*FulfillmentReverseOrderItemQueryDto)
}

// ReleaseFulfillmentReverseOrderItemQueryDto 释放FulfillmentReverseOrderItemQueryDto
func ReleaseFulfillmentReverseOrderItemQueryDto(v *FulfillmentReverseOrderItemQueryDto) {
	v.FulfillmentOrderNo = ""
	v.BizType = 0
	poolFulfillmentReverseOrderItemQueryDto.Put(v)
}
