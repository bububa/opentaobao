package moscm

import (
	"sync"
)

// DeliveryDto 结构体
type DeliveryDto struct {
	// 商品明细
	ShipItems []ShipItemDto `json:"ship_items,omitempty" xml:"ship_items>ship_item_dto,omitempty"`
	// 承运公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 订单号
	OrderNumber string `json:"order_number,omitempty" xml:"order_number,omitempty"`
	// 承运公司编码
	CompanyCode string `json:"company_code,omitempty" xml:"company_code,omitempty"`
	// 出库时间
	OutboundDate string `json:"outbound_date,omitempty" xml:"outbound_date,omitempty"`
	// 运单号
	WaybillNumber string `json:"waybill_number,omitempty" xml:"waybill_number,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
}

var poolDeliveryDto = sync.Pool{
	New: func() any {
		return new(DeliveryDto)
	},
}

// GetDeliveryDto() 从对象池中获取DeliveryDto
func GetDeliveryDto() *DeliveryDto {
	return poolDeliveryDto.Get().(*DeliveryDto)
}

// ReleaseDeliveryDto 释放DeliveryDto
func ReleaseDeliveryDto(v *DeliveryDto) {
	v.ShipItems = v.ShipItems[:0]
	v.CompanyName = ""
	v.OrderNumber = ""
	v.CompanyCode = ""
	v.OutboundDate = ""
	v.WaybillNumber = ""
	v.Remark = ""
	poolDeliveryDto.Put(v)
}
