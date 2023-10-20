package tbk

import (
	"sync"
)

// RegisterInfoDto 结构体
type RegisterInfoDto struct {
	// 渠道独有 -店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 渠道独有 -店铺类型
	ShopType string `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
	// 渠道独有 - 信息
	PhoneNumber string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 渠道独有 -信息
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 渠道独有 -区
	Location string `json:"location,omitempty" xml:"location,omitempty"`
	// 渠道独有 -类型
	ShopCertifyType string `json:"shop_certify_type,omitempty" xml:"shop_certify_type,omitempty"`
	// 渠道独有 - 编号
	CertifyNumber string `json:"certify_number,omitempty" xml:"certify_number,omitempty"`
	// 渠道独有 -经营类型
	Career string `json:"career,omitempty" xml:"career,omitempty"`
}

var poolRegisterInfoDto = sync.Pool{
	New: func() any {
		return new(RegisterInfoDto)
	},
}

// GetRegisterInfoDto() 从对象池中获取RegisterInfoDto
func GetRegisterInfoDto() *RegisterInfoDto {
	return poolRegisterInfoDto.Get().(*RegisterInfoDto)
}

// ReleaseRegisterInfoDto 释放RegisterInfoDto
func ReleaseRegisterInfoDto(v *RegisterInfoDto) {
	v.ShopName = ""
	v.ShopType = ""
	v.PhoneNumber = ""
	v.DetailAddress = ""
	v.Location = ""
	v.ShopCertifyType = ""
	v.CertifyNumber = ""
	v.Career = ""
	poolRegisterInfoDto.Put(v)
}
