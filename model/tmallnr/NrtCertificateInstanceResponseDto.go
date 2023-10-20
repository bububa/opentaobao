package tmallnr

import (
	"sync"
)

// NrtCertificateInstanceResponseDto 结构体
type NrtCertificateInstanceResponseDto struct {
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 电子凭证码值
	CertificateCode string `json:"certificate_code,omitempty" xml:"certificate_code,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 淘系加密ID
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 电子凭证模版id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 投放渠道
	Channel int64 `json:"channel,omitempty" xml:"channel,omitempty"`
}

var poolNrtCertificateInstanceResponseDto = sync.Pool{
	New: func() any {
		return new(NrtCertificateInstanceResponseDto)
	},
}

// GetNrtCertificateInstanceResponseDto() 从对象池中获取NrtCertificateInstanceResponseDto
func GetNrtCertificateInstanceResponseDto() *NrtCertificateInstanceResponseDto {
	return poolNrtCertificateInstanceResponseDto.Get().(*NrtCertificateInstanceResponseDto)
}

// ReleaseNrtCertificateInstanceResponseDto 释放NrtCertificateInstanceResponseDto
func ReleaseNrtCertificateInstanceResponseDto(v *NrtCertificateInstanceResponseDto) {
	v.CreateTime = ""
	v.StartTime = ""
	v.EndTime = ""
	v.CertificateCode = ""
	v.ModifiedTime = ""
	v.OpenId = ""
	v.OutId = ""
	v.Status = 0
	v.Channel = 0
	poolNrtCertificateInstanceResponseDto.Put(v)
}
