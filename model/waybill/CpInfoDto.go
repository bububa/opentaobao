package waybill

import (
	"sync"
)

// CpInfoDto 结构体
type CpInfoDto struct {
	// cp编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// cp名称
	CpName string `json:"cp_name,omitempty" xml:"cp_name,omitempty"`
}

var poolCpInfoDto = sync.Pool{
	New: func() any {
		return new(CpInfoDto)
	},
}

// GetCpInfoDto() 从对象池中获取CpInfoDto
func GetCpInfoDto() *CpInfoDto {
	return poolCpInfoDto.Get().(*CpInfoDto)
}

// ReleaseCpInfoDto 释放CpInfoDto
func ReleaseCpInfoDto(v *CpInfoDto) {
	v.CpCode = ""
	v.CpName = ""
	poolCpInfoDto.Put(v)
}
