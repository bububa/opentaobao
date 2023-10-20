package logistic

import (
	"sync"
)

// WmsMaterialPackageDto 结构体
type WmsMaterialPackageDto struct {
	// 包材信息
	Materials []WmsMaterialDetailDto `json:"materials,omitempty" xml:"materials>wms_material_detail_dto,omitempty"`
	// 商品信息
	Items []WmsMaterialPackageItemDto `json:"items,omitempty" xml:"items>wms_material_package_item_dto,omitempty"`
	// 包裹编码
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// SF23434
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
	// SF
	LogisticsCode string `json:"logistics_code,omitempty" xml:"logistics_code,omitempty"`
}

var poolWmsMaterialPackageDto = sync.Pool{
	New: func() any {
		return new(WmsMaterialPackageDto)
	},
}

// GetWmsMaterialPackageDto() 从对象池中获取WmsMaterialPackageDto
func GetWmsMaterialPackageDto() *WmsMaterialPackageDto {
	return poolWmsMaterialPackageDto.Get().(*WmsMaterialPackageDto)
}

// ReleaseWmsMaterialPackageDto 释放WmsMaterialPackageDto
func ReleaseWmsMaterialPackageDto(v *WmsMaterialPackageDto) {
	v.Materials = v.Materials[:0]
	v.Items = v.Items[:0]
	v.PackageCode = ""
	v.ExpressCode = ""
	v.LogisticsCode = ""
	poolWmsMaterialPackageDto.Put(v)
}
