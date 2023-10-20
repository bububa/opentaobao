package lstvending

import (
	"sync"
)

// OpenEquipmentDto 结构体
type OpenEquipmentDto struct {
	// 省份代码
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 城市代码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 地区代码
	AreaCode string `json:"area_code,omitempty" xml:"area_code,omitempty"`
	// 地区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 供应商代码
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 型号名称
	ModelName string `json:"model_name,omitempty" xml:"model_name,omitempty"`
	// 设备代码
	EquipmentCode string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
	// 设备激活时间
	EnabledTime string `json:"enabled_time,omitempty" xml:"enabled_time,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 设备回收（转租）时间
	RecoveredTime string `json:"recovered_time,omitempty" xml:"recovered_time,omitempty"`
	// 设备ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 创建时间
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 设备状态：1未激活，2已激活，3已回收（转租）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolOpenEquipmentDto = sync.Pool{
	New: func() any {
		return new(OpenEquipmentDto)
	},
}

// GetOpenEquipmentDto() 从对象池中获取OpenEquipmentDto
func GetOpenEquipmentDto() *OpenEquipmentDto {
	return poolOpenEquipmentDto.Get().(*OpenEquipmentDto)
}

// ReleaseOpenEquipmentDto 释放OpenEquipmentDto
func ReleaseOpenEquipmentDto(v *OpenEquipmentDto) {
	v.ProvinceCode = ""
	v.Province = ""
	v.CityCode = ""
	v.City = ""
	v.AreaCode = ""
	v.Area = ""
	v.SupplierCode = ""
	v.ModelName = ""
	v.EquipmentCode = ""
	v.EnabledTime = ""
	v.GmtModified = ""
	v.RecoveredTime = ""
	v.Id = 0
	v.GmtCreate = 0
	v.Status = 0
	poolOpenEquipmentDto.Put(v)
}
