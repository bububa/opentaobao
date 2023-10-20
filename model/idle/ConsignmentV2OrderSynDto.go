package idle

import (
	"sync"
)

// ConsignmentV2OrderSynDto 结构体
type ConsignmentV2OrderSynDto struct {
	// 订单子状态
	OrderSubStatus string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
	// 订单主状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 闲鱼订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 下单环境
	Env string `json:"env,omitempty" xml:"env,omitempty"`
	// 不同的状态传递不同参数
	Attribute *Attribute `json:"attribute,omitempty" xml:"attribute,omitempty"`
}

var poolConsignmentV2OrderSynDto = sync.Pool{
	New: func() any {
		return new(ConsignmentV2OrderSynDto)
	},
}

// GetConsignmentV2OrderSynDto() 从对象池中获取ConsignmentV2OrderSynDto
func GetConsignmentV2OrderSynDto() *ConsignmentV2OrderSynDto {
	return poolConsignmentV2OrderSynDto.Get().(*ConsignmentV2OrderSynDto)
}

// ReleaseConsignmentV2OrderSynDto 释放ConsignmentV2OrderSynDto
func ReleaseConsignmentV2OrderSynDto(v *ConsignmentV2OrderSynDto) {
	v.OrderSubStatus = ""
	v.OrderStatus = ""
	v.BizOrderId = ""
	v.Env = ""
	v.Attribute = nil
	poolConsignmentV2OrderSynDto.Put(v)
}
