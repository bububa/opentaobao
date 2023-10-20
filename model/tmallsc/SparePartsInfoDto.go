package tmallsc

import (
	"sync"
)

// SparePartsInfoDto 结构体
type SparePartsInfoDto struct {
	// 备件编码
	SparePartsCode string `json:"spare_parts_code,omitempty" xml:"spare_parts_code,omitempty"`
	// 备件名称
	SparePartsName string `json:"spare_parts_name,omitempty" xml:"spare_parts_name,omitempty"`
	// 备件型号
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 备件类型
	SparePartsType string `json:"spare_parts_type,omitempty" xml:"spare_parts_type,omitempty"`
	// 备件一级类目
	FirstLevelCategoryId int64 `json:"first_level_category_id,omitempty" xml:"first_level_category_id,omitempty"`
	// 备件销售价
	SellingPrice int64 `json:"selling_price,omitempty" xml:"selling_price,omitempty"`
	// 备件二级类目
	SecondLevelCategoryId int64 `json:"second_level_category_id,omitempty" xml:"second_level_category_id,omitempty"`
	// 备件来源
	SparePartsSource int64 `json:"spare_parts_source,omitempty" xml:"spare_parts_source,omitempty"`
	// 备件进价
	PurchasePrice int64 `json:"purchase_price,omitempty" xml:"purchase_price,omitempty"`
	// 备件数量
	SparePartsCount int64 `json:"spare_parts_count,omitempty" xml:"spare_parts_count,omitempty"`
	// 是否需要返件
	NeedReturn bool `json:"need_return,omitempty" xml:"need_return,omitempty"`
}

var poolSparePartsInfoDto = sync.Pool{
	New: func() any {
		return new(SparePartsInfoDto)
	},
}

// GetSparePartsInfoDto() 从对象池中获取SparePartsInfoDto
func GetSparePartsInfoDto() *SparePartsInfoDto {
	return poolSparePartsInfoDto.Get().(*SparePartsInfoDto)
}

// ReleaseSparePartsInfoDto 释放SparePartsInfoDto
func ReleaseSparePartsInfoDto(v *SparePartsInfoDto) {
	v.SparePartsCode = ""
	v.SparePartsName = ""
	v.Model = ""
	v.SparePartsType = ""
	v.FirstLevelCategoryId = 0
	v.SellingPrice = 0
	v.SecondLevelCategoryId = 0
	v.SparePartsSource = 0
	v.PurchasePrice = 0
	v.SparePartsCount = 0
	v.NeedReturn = false
	poolSparePartsInfoDto.Put(v)
}
