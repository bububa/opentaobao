package aliexpresssumaitong

import (
	"sync"
)

// PlatformParameterDto 结构体
type PlatformParameterDto struct {
	// 发货人/平台公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 发货人/平台联系人姓名
	ContactorName string `json:"contactor_name,omitempty" xml:"contactor_name,omitempty"`
	// 发货人/平台国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 发货人/平台详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 发货人/平台城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// ISO国家编码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 发货人/平台电话
	ContactNumber string `json:"contact_number,omitempty" xml:"contact_number,omitempty"`
	// 发货人/平台邮箱
	EmailAddress string `json:"email_address,omitempty" xml:"email_address,omitempty"`
	// IOSS 平台税号
	IossNumber string `json:"ioss_number,omitempty" xml:"ioss_number,omitempty"`
}

var poolPlatformParameterDto = sync.Pool{
	New: func() any {
		return new(PlatformParameterDto)
	},
}

// GetPlatformParameterDto() 从对象池中获取PlatformParameterDto
func GetPlatformParameterDto() *PlatformParameterDto {
	return poolPlatformParameterDto.Get().(*PlatformParameterDto)
}

// ReleasePlatformParameterDto 释放PlatformParameterDto
func ReleasePlatformParameterDto(v *PlatformParameterDto) {
	v.CompanyName = ""
	v.ContactorName = ""
	v.Country = ""
	v.Address = ""
	v.City = ""
	v.CountryCode = ""
	v.ContactNumber = ""
	v.EmailAddress = ""
	v.IossNumber = ""
	poolPlatformParameterDto.Put(v)
}
