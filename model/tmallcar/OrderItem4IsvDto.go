package tmallcar

import (
	"sync"
)

// OrderItem4IsvDto 结构体
type OrderItem4IsvDto struct {
	// 实际支付金额
	ActualTotalFee string `json:"actual_total_fee,omitempty" xml:"actual_total_fee,omitempty"`
	// 扩展属性
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 商品头图
	ItemHeadImg string `json:"item_head_img,omitempty" xml:"item_head_img,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 订单商品原价
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 是否在三包期内
	InGuarantee bool `json:"in_guarantee,omitempty" xml:"in_guarantee,omitempty"`
}

var poolOrderItem4IsvDto = sync.Pool{
	New: func() any {
		return new(OrderItem4IsvDto)
	},
}

// GetOrderItem4IsvDto() 从对象池中获取OrderItem4IsvDto
func GetOrderItem4IsvDto() *OrderItem4IsvDto {
	return poolOrderItem4IsvDto.Get().(*OrderItem4IsvDto)
}

// ReleaseOrderItem4IsvDto 释放OrderItem4IsvDto
func ReleaseOrderItem4IsvDto(v *OrderItem4IsvDto) {
	v.ActualTotalFee = ""
	v.Extension = ""
	v.ItemHeadImg = ""
	v.ItemName = ""
	v.TotalFee = ""
	v.ItemId = 0
	v.InGuarantee = false
	poolOrderItem4IsvDto.Put(v)
}
