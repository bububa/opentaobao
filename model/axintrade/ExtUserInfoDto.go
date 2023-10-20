package axintrade

import (
	"sync"
)

// ExtUserInfoDto 结构体
type ExtUserInfoDto struct {
	// 指定买家支付宝实名证件号
	ExtUserCertNo string `json:"ext_user_cert_no,omitempty" xml:"ext_user_cert_no,omitempty"`
	// 指定买家支付宝实名证件类型
	ExtUserCertType string `json:"ext_user_cert_type,omitempty" xml:"ext_user_cert_type,omitempty"`
	// 指定买家支付宝实名姓名
	ExtUserName string `json:"ext_user_name,omitempty" xml:"ext_user_name,omitempty"`
}

var poolExtUserInfoDto = sync.Pool{
	New: func() any {
		return new(ExtUserInfoDto)
	},
}

// GetExtUserInfoDto() 从对象池中获取ExtUserInfoDto
func GetExtUserInfoDto() *ExtUserInfoDto {
	return poolExtUserInfoDto.Get().(*ExtUserInfoDto)
}

// ReleaseExtUserInfoDto 释放ExtUserInfoDto
func ReleaseExtUserInfoDto(v *ExtUserInfoDto) {
	v.ExtUserCertNo = ""
	v.ExtUserCertType = ""
	v.ExtUserName = ""
	poolExtUserInfoDto.Put(v)
}
