package xiamicontent

import (
	"sync"
)

// CompanyDto 结构体
type CompanyDto struct {
	// 厂牌名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 厂牌id
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
}

var poolCompanyDto = sync.Pool{
	New: func() any {
		return new(CompanyDto)
	},
}

// GetCompanyDto() 从对象池中获取CompanyDto
func GetCompanyDto() *CompanyDto {
	return poolCompanyDto.Get().(*CompanyDto)
}

// ReleaseCompanyDto 释放CompanyDto
func ReleaseCompanyDto(v *CompanyDto) {
	v.CompanyName = ""
	v.CompanyId = 0
	poolCompanyDto.Put(v)
}
