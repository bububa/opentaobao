package alihealth2

import (
	"sync"
)

// DrugInfoDto 结构体
type DrugInfoDto struct {
	// 药品过期日期
	ExpiryDate string `json:"expiry_date,omitempty" xml:"expiry_date,omitempty"`
	// 药品通用名
	DrugName string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
	// 生产批号
	ProductionBatch string `json:"production_batch,omitempty" xml:"production_batch,omitempty"`
	// 药监码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 规格
	Specifications string `json:"specifications,omitempty" xml:"specifications,omitempty"`
	// 包装
	PkgSpec string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
	// 剂型
	PrepnType string `json:"prepn_type,omitempty" xml:"prepn_type,omitempty"`
	// 生产企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
}

var poolDrugInfoDto = sync.Pool{
	New: func() any {
		return new(DrugInfoDto)
	},
}

// GetDrugInfoDto() 从对象池中获取DrugInfoDto
func GetDrugInfoDto() *DrugInfoDto {
	return poolDrugInfoDto.Get().(*DrugInfoDto)
}

// ReleaseDrugInfoDto 释放DrugInfoDto
func ReleaseDrugInfoDto(v *DrugInfoDto) {
	v.ExpiryDate = ""
	v.DrugName = ""
	v.ProductionBatch = ""
	v.Code = ""
	v.Specifications = ""
	v.PkgSpec = ""
	v.PrepnType = ""
	v.EntName = ""
	poolDrugInfoDto.Put(v)
}
