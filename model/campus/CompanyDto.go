package campus

import (
	"sync"
)

// CompanyDto 结构体
type CompanyDto struct {
	// 公司编号
	CompanyCode string `json:"company_code,omitempty" xml:"company_code,omitempty"`
	// corpId
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 签约主体公司id
	HrSignCompanyId string `json:"hr_sign_company_id,omitempty" xml:"hr_sign_company_id,omitempty"`
	// 园区名称
	CampusName string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
	// 公司联系电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 公司状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 公司简称
	ShortName string `json:"short_name,omitempty" xml:"short_name,omitempty"`
	// 公司名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 公司人数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 园区ID
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 公司ID
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 公司ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 是否物业公司
	IsWuye bool `json:"is_wuye,omitempty" xml:"is_wuye,omitempty"`
	// 是否默认公司
	IsDefault bool `json:"is_default,omitempty" xml:"is_default,omitempty"`
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
	v.CompanyCode = ""
	v.CorpId = ""
	v.HrSignCompanyId = ""
	v.CampusName = ""
	v.Mobile = ""
	v.Status = ""
	v.ShortName = ""
	v.Name = ""
	v.Count = 0
	v.CampusId = 0
	v.CompanyId = 0
	v.Id = 0
	v.IsWuye = false
	v.IsDefault = false
	poolCompanyDto.Put(v)
}
