package lstvending

import (
	"sync"
)

// VendingCargoSpaceDto 结构体
type VendingCargoSpaceDto struct {
	// 外部商品ID
	ExternalGoodsId string `json:"external_goods_id,omitempty" xml:"external_goods_id,omitempty"`
	// 厂商设备唯一编码
	EquipmentCode string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
	// 厂商货道ID
	ExternalId string `json:"external_id,omitempty" xml:"external_id,omitempty"`
	// 供应商编码
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 修改时间
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 货道商品数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 创建时间
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 商品折扣价
	DiscountPrice int64 `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	// 货道编号，从1开始
	CargoRoadNo int64 `json:"cargo_road_no,omitempty" xml:"cargo_road_no,omitempty"`
	// 货道类型：1普通货道，2VIP货道
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 货架编号，从1开始
	ShelfNo int64 `json:"shelf_no,omitempty" xml:"shelf_no,omitempty"`
	// 商品单价，单位：分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 货道状态：1正常，2故障
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 商品ID
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 货道ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolVendingCargoSpaceDto = sync.Pool{
	New: func() any {
		return new(VendingCargoSpaceDto)
	},
}

// GetVendingCargoSpaceDto() 从对象池中获取VendingCargoSpaceDto
func GetVendingCargoSpaceDto() *VendingCargoSpaceDto {
	return poolVendingCargoSpaceDto.Get().(*VendingCargoSpaceDto)
}

// ReleaseVendingCargoSpaceDto 释放VendingCargoSpaceDto
func ReleaseVendingCargoSpaceDto(v *VendingCargoSpaceDto) {
	v.ExternalGoodsId = ""
	v.EquipmentCode = ""
	v.ExternalId = ""
	v.SupplierCode = ""
	v.GmtModified = 0
	v.Count = 0
	v.GmtCreate = 0
	v.DiscountPrice = 0
	v.CargoRoadNo = 0
	v.BizType = 0
	v.ShelfNo = 0
	v.Price = 0
	v.Status = 0
	v.GoodsId = 0
	v.Id = 0
	poolVendingCargoSpaceDto.Put(v)
}
