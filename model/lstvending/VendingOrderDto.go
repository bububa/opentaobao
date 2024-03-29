package lstvending

import (
	"sync"
)

// VendingOrderDto 结构体
type VendingOrderDto struct {
	// 发货清单
	ShippedEquipmentList []VendingShippedEquipmentDto `json:"shipped_equipment_list,omitempty" xml:"shipped_equipment_list>vending_shipped_equipment_dto,omitempty"`
	// 快递单号
	TrackingNo string `json:"tracking_no,omitempty" xml:"tracking_no,omitempty"`
	// 快递联系人
	ShippingContact string `json:"shipping_contact,omitempty" xml:"shipping_contact,omitempty"`
	// 供应商订单号
	IsvOrderNo string `json:"isv_order_no,omitempty" xml:"isv_order_no,omitempty"`
	// 快递联系电话
	ShippingContactTel string `json:"shipping_contact_tel,omitempty" xml:"shipping_contact_tel,omitempty"`
	// 修改时间
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 零售通订单号
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 创建时间
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 发货时间
	DeliveryTime int64 `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
}

var poolVendingOrderDto = sync.Pool{
	New: func() any {
		return new(VendingOrderDto)
	},
}

// GetVendingOrderDto() 从对象池中获取VendingOrderDto
func GetVendingOrderDto() *VendingOrderDto {
	return poolVendingOrderDto.Get().(*VendingOrderDto)
}

// ReleaseVendingOrderDto 释放VendingOrderDto
func ReleaseVendingOrderDto(v *VendingOrderDto) {
	v.ShippedEquipmentList = v.ShippedEquipmentList[:0]
	v.TrackingNo = ""
	v.ShippingContact = ""
	v.IsvOrderNo = ""
	v.ShippingContactTel = ""
	v.GmtModified = 0
	v.Id = 0
	v.GmtCreate = 0
	v.DeliveryTime = 0
	poolVendingOrderDto.Put(v)
}
