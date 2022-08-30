package wlbimports

// ContactInfoRequest 结构体
type ContactInfoRequest struct {
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 公司
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 详细地址（最大135个字符）
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 地区id
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 国家id
	CountryId int64 `json:"country_id,omitempty" xml:"country_id,omitempty"`
}
