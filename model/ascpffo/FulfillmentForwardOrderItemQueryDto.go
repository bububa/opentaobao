package ascpffo

import (
	"sync"
)

// FulfillmentForwardOrderItemQueryDto 结构体
type FulfillmentForwardOrderItemQueryDto struct {
	// 履约单号
	FulfillmentOrderNo string `json:"fulfillment_order_no,omitempty" xml:"fulfillment_order_no,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

var poolFulfillmentForwardOrderItemQueryDto = sync.Pool{
	New: func() any {
		return new(FulfillmentForwardOrderItemQueryDto)
	},
}

// GetFulfillmentForwardOrderItemQueryDto() 从对象池中获取FulfillmentForwardOrderItemQueryDto
func GetFulfillmentForwardOrderItemQueryDto() *FulfillmentForwardOrderItemQueryDto {
	return poolFulfillmentForwardOrderItemQueryDto.Get().(*FulfillmentForwardOrderItemQueryDto)
}

// ReleaseFulfillmentForwardOrderItemQueryDto 释放FulfillmentForwardOrderItemQueryDto
func ReleaseFulfillmentForwardOrderItemQueryDto(v *FulfillmentForwardOrderItemQueryDto) {
	v.FulfillmentOrderNo = ""
	v.BizType = 0
	poolFulfillmentForwardOrderItemQueryDto.Put(v)
}
