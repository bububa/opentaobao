package logistic

import (
	"sync"
)

// WmsMaterialPackageItemDto 结构体
type WmsMaterialPackageItemDto struct {
	// 耗材信息
	Materials []WmsMaterialDetailDto `json:"materials,omitempty" xml:"materials>wms_material_detail_dto,omitempty"`
	// 商品仓储系统编码
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货主编码
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
}

var poolWmsMaterialPackageItemDto = sync.Pool{
	New: func() any {
		return new(WmsMaterialPackageItemDto)
	},
}

// GetWmsMaterialPackageItemDto() 从对象池中获取WmsMaterialPackageItemDto
func GetWmsMaterialPackageItemDto() *WmsMaterialPackageItemDto {
	return poolWmsMaterialPackageItemDto.Get().(*WmsMaterialPackageItemDto)
}

// ReleaseWmsMaterialPackageItemDto 释放WmsMaterialPackageItemDto
func ReleaseWmsMaterialPackageItemDto(v *WmsMaterialPackageItemDto) {
	v.Materials = v.Materials[:0]
	v.ItemId = ""
	v.OwnerCode = ""
	v.ItemCode = ""
	poolWmsMaterialPackageItemDto.Put(v)
}
