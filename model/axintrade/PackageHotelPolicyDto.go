package axintrade

import (
	"sync"
)

// PackageHotelPolicyDto 结构体
type PackageHotelPolicyDto struct {
	// 入住方式
	CheckInTypes []string `json:"check_in_types,omitempty" xml:"check_in_types>string,omitempty"`
	// 入住时间
	CheckInTime string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// 离店时间
	CheckOutTime string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// 入住方式其他说明
	CheckInTypeRemark string `json:"check_in_type_remark,omitempty" xml:"check_in_type_remark,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
}

var poolPackageHotelPolicyDto = sync.Pool{
	New: func() any {
		return new(PackageHotelPolicyDto)
	},
}

// GetPackageHotelPolicyDto() 从对象池中获取PackageHotelPolicyDto
func GetPackageHotelPolicyDto() *PackageHotelPolicyDto {
	return poolPackageHotelPolicyDto.Get().(*PackageHotelPolicyDto)
}

// ReleasePackageHotelPolicyDto 释放PackageHotelPolicyDto
func ReleasePackageHotelPolicyDto(v *PackageHotelPolicyDto) {
	v.CheckInTypes = v.CheckInTypes[:0]
	v.CheckInTime = ""
	v.CheckOutTime = ""
	v.CheckInTypeRemark = ""
	v.Remark = ""
	poolPackageHotelPolicyDto.Put(v)
}
