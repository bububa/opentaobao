package idle

import (
	"sync"
)

// TenderOrderSynDto 结构体
type TenderOrderSynDto struct {
	// 订单子状态
	OrderSubStatus string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
	// 订单id
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 流量来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 供商家传入的订单履约数据
	Attribute *HashMap `json:"attribute,omitempty" xml:"attribute,omitempty"`
}

var poolTenderOrderSynDto = sync.Pool{
	New: func() any {
		return new(TenderOrderSynDto)
	},
}

// GetTenderOrderSynDto() 从对象池中获取TenderOrderSynDto
func GetTenderOrderSynDto() *TenderOrderSynDto {
	return poolTenderOrderSynDto.Get().(*TenderOrderSynDto)
}

// ReleaseTenderOrderSynDto 释放TenderOrderSynDto
func ReleaseTenderOrderSynDto(v *TenderOrderSynDto) {
	v.OrderSubStatus = ""
	v.BizOrderId = ""
	v.OrderStatus = ""
	v.Source = ""
	v.Attribute = nil
	poolTenderOrderSynDto.Put(v)
}
