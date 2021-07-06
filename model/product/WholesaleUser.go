package product

// WholesaleUser 结构体
type WholesaleUser struct {
	// user id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// user's company address用户公司地址
	CompAddress string `json:"comp_address,omitempty" xml:"comp_address,omitempty"`
	// company country公司所在国家
	CompCountry string `json:"comp_country,omitempty" xml:"comp_country,omitempty"`
	// company city公司所在城市
	CompCity string `json:"comp_city,omitempty" xml:"comp_city,omitempty"`
	// company province 公司所在的省
	CompProvince string `json:"comp_province,omitempty" xml:"comp_province,omitempty"`
	// company logo url公司logo 地址
	LogoUrl string `json:"logo_url,omitempty" xml:"logo_url,omitempty"`
	// company name公司名
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 公司所在地邮政编码
	CompZip string `json:"comp_zip,omitempty" xml:"comp_zip,omitempty"`
	// golden_supplier_years公司成为金牌会员的年份
	GoldenSupplierYears int64 `json:"golden_supplier_years,omitempty" xml:"golden_supplier_years,omitempty"`
}
