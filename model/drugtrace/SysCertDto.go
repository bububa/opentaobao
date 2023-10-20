package drugtrace

import (
	"sync"
)

// SysCertDto 结构体
type SysCertDto struct {
	// 证书序列号
	CertSn string `json:"cert_sn,omitempty" xml:"cert_sn,omitempty"`
	// 证书公钥
	Cert string `json:"cert,omitempty" xml:"cert,omitempty"`
}

var poolSysCertDto = sync.Pool{
	New: func() any {
		return new(SysCertDto)
	},
}

// GetSysCertDto() 从对象池中获取SysCertDto
func GetSysCertDto() *SysCertDto {
	return poolSysCertDto.Get().(*SysCertDto)
}

// ReleaseSysCertDto 释放SysCertDto
func ReleaseSysCertDto(v *SysCertDto) {
	v.CertSn = ""
	v.Cert = ""
	poolSysCertDto.Put(v)
}
