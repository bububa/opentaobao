package axintrade

import (
	"sync"
)

// PackageCateringInfoDto 结构体
type PackageCateringInfoDto struct {
	// 餐饮名称
	CateringName string `json:"catering_name,omitempty" xml:"catering_name,omitempty"`
	// 餐饮份数为其他时，其他说明
	CateringRemark string `json:"catering_remark,omitempty" xml:"catering_remark,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 餐饮类型
	CateringType int64 `json:"catering_type,omitempty" xml:"catering_type,omitempty"`
	// 餐饮份数类型
	CateringNumType int64 `json:"catering_num_type,omitempty" xml:"catering_num_type,omitempty"`
	// 餐饮份数
	CateringNum int64 `json:"catering_num,omitempty" xml:"catering_num,omitempty"`
}

var poolPackageCateringInfoDto = sync.Pool{
	New: func() any {
		return new(PackageCateringInfoDto)
	},
}

// GetPackageCateringInfoDto() 从对象池中获取PackageCateringInfoDto
func GetPackageCateringInfoDto() *PackageCateringInfoDto {
	return poolPackageCateringInfoDto.Get().(*PackageCateringInfoDto)
}

// ReleasePackageCateringInfoDto 释放PackageCateringInfoDto
func ReleasePackageCateringInfoDto(v *PackageCateringInfoDto) {
	v.CateringName = ""
	v.CateringRemark = ""
	v.Remark = ""
	v.CateringType = 0
	v.CateringNumType = 0
	v.CateringNum = 0
	poolPackageCateringInfoDto.Put(v)
}
