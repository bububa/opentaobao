package alihouse

import (
	"sync"
)

// CompanyBrandCityDto 结构体
type CompanyBrandCityDto struct {
	// 外部城市品牌店ID
	OuterCompanyBrandId string `json:"outer_company_brand_id,omitempty" xml:"outer_company_brand_id,omitempty"`
	// 公司外部ID
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 外部品牌ID
	OuterBrandId string `json:"outer_brand_id,omitempty" xml:"outer_brand_id,omitempty"`
	// 公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 公司简称
	CompanyNameShort string `json:"company_name_short,omitempty" xml:"company_name_short,omitempty"`
	// 营业执照编号
	CompanyLicenseNo string `json:"company_license_no,omitempty" xml:"company_license_no,omitempty"`
	// 企业法定代表人
	CompanyLegalPerson string `json:"company_legal_person,omitempty" xml:"company_legal_person,omitempty"`
	// 营业执照图片
	CompanyLicensePhoto string `json:"company_license_photo,omitempty" xml:"company_license_photo,omitempty"`
	// 营业执照有效时间
	CompanyLicenseExpireTime string `json:"company_license_expire_time,omitempty" xml:"company_license_expire_time,omitempty"`
	// 公司logo
	CompanyLogo string `json:"company_logo,omitempty" xml:"company_logo,omitempty"`
	// 备案证明编号
	CompanyCertificateNo string `json:"company_certificate_no,omitempty" xml:"company_certificate_no,omitempty"`
	// 备案证明图片
	CompanyCertificatePhoto string `json:"company_certificate_photo,omitempty" xml:"company_certificate_photo,omitempty"`
	// 联系人
	ContactMan string `json:"contact_man,omitempty" xml:"contact_man,omitempty"`
	// 联系电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 高德经度
	GaodeLongitude string `json:"gaode_longitude,omitempty" xml:"gaode_longitude,omitempty"`
	// 高德纬度
	GaodeLatitude string `json:"gaode_latitude,omitempty" xml:"gaode_latitude,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 品牌门店描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 城市品牌店标签
	TagCodes string `json:"tag_codes,omitempty" xml:"tag_codes,omitempty"`
	// 城市ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 是否长期有效 1-是 0-否
	CompanyLicenseStatus int64 `json:"company_license_status,omitempty" xml:"company_license_status,omitempty"`
	// 是否删除 0 - 否 1 - 是
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 是否虚拟品牌店 1-是 0-否
	IsVirtualCompany int64 `json:"is_virtual_company,omitempty" xml:"is_virtual_company,omitempty"`
	// 是否为测试标 0-否 1-是
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 城市品牌类型：交易服务类型-5、二租业务类型-0
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolCompanyBrandCityDto = sync.Pool{
	New: func() any {
		return new(CompanyBrandCityDto)
	},
}

// GetCompanyBrandCityDto() 从对象池中获取CompanyBrandCityDto
func GetCompanyBrandCityDto() *CompanyBrandCityDto {
	return poolCompanyBrandCityDto.Get().(*CompanyBrandCityDto)
}

// ReleaseCompanyBrandCityDto 释放CompanyBrandCityDto
func ReleaseCompanyBrandCityDto(v *CompanyBrandCityDto) {
	v.OuterCompanyBrandId = ""
	v.OuterCompanyId = ""
	v.OuterBrandId = ""
	v.CompanyName = ""
	v.CompanyNameShort = ""
	v.CompanyLicenseNo = ""
	v.CompanyLegalPerson = ""
	v.CompanyLicensePhoto = ""
	v.CompanyLicenseExpireTime = ""
	v.CompanyLogo = ""
	v.CompanyCertificateNo = ""
	v.CompanyCertificatePhoto = ""
	v.ContactMan = ""
	v.ContactPhone = ""
	v.GaodeLongitude = ""
	v.GaodeLatitude = ""
	v.Address = ""
	v.Description = ""
	v.TagCodes = ""
	v.CityId = 0
	v.CompanyLicenseStatus = 0
	v.IsDeleted = 0
	v.IsVirtualCompany = 0
	v.IsTest = 0
	v.Type = 0
	poolCompanyBrandCityDto.Put(v)
}
