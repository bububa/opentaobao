package moscm

import (
	"sync"
)

// RmaItemDto 结构体
type RmaItemDto struct {
	// 商品类型,可选值：普通商品(NORMAL),赠品(GIFT)
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 退/换货原因详情
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 退/换货原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 折扣金额，单位:元
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 售价，单位:元
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 商品属性
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 外部商品Id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 商品Id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 实付金额，单位:元
	ActualAmount string `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	// 唯一码（仅限西有使用）
	DiscCode string `json:"disc_code,omitempty" xml:"disc_code,omitempty"`
	// 商品编码
	SettlementCode string `json:"settlement_code,omitempty" xml:"settlement_code,omitempty"`
}

var poolRmaItemDto = sync.Pool{
	New: func() any {
		return new(RmaItemDto)
	},
}

// GetRmaItemDto() 从对象池中获取RmaItemDto
func GetRmaItemDto() *RmaItemDto {
	return poolRmaItemDto.Get().(*RmaItemDto)
}

// ReleaseRmaItemDto 释放RmaItemDto
func ReleaseRmaItemDto(v *RmaItemDto) {
	v.ProductType = ""
	v.Desc = ""
	v.Reason = ""
	v.Discount = ""
	v.Price = ""
	v.Quantity = ""
	v.Properties = ""
	v.Title = ""
	v.OutId = ""
	v.SkuId = ""
	v.ActualAmount = ""
	v.DiscCode = ""
	v.SettlementCode = ""
	poolRmaItemDto.Put(v)
}
