package alihouse

// CompanyDto 结构体
type CompanyDto struct {
	// 公司logo
	CompanyLogo string `json:"company_logo,omitempty" xml:"company_logo,omitempty"`
	// 营业执照过期时间
	CompanyLicenseExpireTime string `json:"company_license_expire_time,omitempty" xml:"company_license_expire_time,omitempty"`
	// 营业执照图片
	CompanyLicensePhoto string `json:"company_license_photo,omitempty" xml:"company_license_photo,omitempty"`
	// 法定代表人
	CompanyLegalPerson string `json:"company_legal_person,omitempty" xml:"company_legal_person,omitempty"`
	// 营业执照编号
	CompanyLicenseNo string `json:"company_license_no,omitempty" xml:"company_license_no,omitempty"`
	// 公司简称
	CompanyNameShort string `json:"company_name_short,omitempty" xml:"company_name_short,omitempty"`
	// 公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 外部公司ID
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 高德纬度
	GaodeLatitude string `json:"gaode_latitude,omitempty" xml:"gaode_latitude,omitempty"`
	// 高德经度
	GaodeLongitude string `json:"gaode_longitude,omitempty" xml:"gaode_longitude,omitempty"`
	// 联系电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 联系人
	ContactMan string `json:"contact_man,omitempty" xml:"contact_man,omitempty"`
	// 备案证明图片
	CompanyCertificatePhoto string `json:"company_certificate_photo,omitempty" xml:"company_certificate_photo,omitempty"`
	// 备案证明编号
	CompanyCertificateNo string `json:"company_certificate_no,omitempty" xml:"company_certificate_no,omitempty"`
	// 是否长期有效 1-是 0-否
	CompanyLicenseStatus int64 `json:"company_license_status,omitempty" xml:"company_license_status,omitempty"`
	// 城市ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 是否删除 1-是 0-否
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 是否测试 0-否 1-是
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
