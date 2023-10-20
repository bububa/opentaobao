package tmallsc

import (
	"sync"
)

// ApplicationInfoDto 结构体
type ApplicationInfoDto struct {
	// 品牌授权商openid
	BrandLicensorOpenId string `json:"brand_licensor_open_id,omitempty" xml:"brand_licensor_open_id,omitempty"`
	// 服务商openid
	SupplierOpenId string `json:"supplier_open_id,omitempty" xml:"supplier_open_id,omitempty"`
	// 备件出库时间
	SparePartsOutTime string `json:"spare_parts_out_time,omitempty" xml:"spare_parts_out_time,omitempty"`
	// 备件申请单id
	SparePartsApplicationId int64 `json:"spare_parts_application_id,omitempty" xml:"spare_parts_application_id,omitempty"`
}

var poolApplicationInfoDto = sync.Pool{
	New: func() any {
		return new(ApplicationInfoDto)
	},
}

// GetApplicationInfoDto() 从对象池中获取ApplicationInfoDto
func GetApplicationInfoDto() *ApplicationInfoDto {
	return poolApplicationInfoDto.Get().(*ApplicationInfoDto)
}

// ReleaseApplicationInfoDto 释放ApplicationInfoDto
func ReleaseApplicationInfoDto(v *ApplicationInfoDto) {
	v.BrandLicensorOpenId = ""
	v.SupplierOpenId = ""
	v.SparePartsOutTime = ""
	v.SparePartsApplicationId = 0
	poolApplicationInfoDto.Put(v)
}
