package omniorder

import (
	"sync"
)

// ScbRefundAddressDto 结构体
type ScbRefundAddressDto struct {
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// 区名称
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 完整地址
	CompleteAddress string `json:"complete_address,omitempty" xml:"complete_address,omitempty"`
	// 收件人
	ConsigneeFullName string `json:"consignee_full_name,omitempty" xml:"consignee_full_name,omitempty"`
	// 固定电话
	FixedPhone string `json:"fixed_phone,omitempty" xml:"fixed_phone,omitempty"`
	// 移动电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 邮编
	PostCode string `json:"post_code,omitempty" xml:"post_code,omitempty"`
	// 省名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 地址id
	ContactId int64 `json:"contact_id,omitempty" xml:"contact_id,omitempty"`
	// 区code
	DivisionCode int64 `json:"division_code,omitempty" xml:"division_code,omitempty"`
}

var poolScbRefundAddressDto = sync.Pool{
	New: func() any {
		return new(ScbRefundAddressDto)
	},
}

// GetScbRefundAddressDto() 从对象池中获取ScbRefundAddressDto
func GetScbRefundAddressDto() *ScbRefundAddressDto {
	return poolScbRefundAddressDto.Get().(*ScbRefundAddressDto)
}

// ReleaseScbRefundAddressDto 释放ScbRefundAddressDto
func ReleaseScbRefundAddressDto(v *ScbRefundAddressDto) {
	v.AddressDetail = ""
	v.AreaName = ""
	v.CityName = ""
	v.CompleteAddress = ""
	v.ConsigneeFullName = ""
	v.FixedPhone = ""
	v.Mobile = ""
	v.PostCode = ""
	v.ProvinceName = ""
	v.ContactId = 0
	v.DivisionCode = 0
	poolScbRefundAddressDto.Put(v)
}
