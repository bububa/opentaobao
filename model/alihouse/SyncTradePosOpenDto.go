package alihouse

// SyncTradePosOpenDto 结构体
type SyncTradePosOpenDto struct {
	// 经营主体id
	MerchantOpenId string `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 国籍
	LegalNationality string `json:"legal_nationality,omitempty" xml:"legal_nationality,omitempty"`
	// 法人联系方式
	LegalPhone string `json:"legal_phone,omitempty" xml:"legal_phone,omitempty"`
	// 法人地址
	LegalAddress string `json:"legal_address,omitempty" xml:"legal_address,omitempty"`
	// 法人性别
	LegalSex string `json:"legal_sex,omitempty" xml:"legal_sex,omitempty"`
	// 省编码
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 市编码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 区编码
	DistrictCode string `json:"district_code,omitempty" xml:"district_code,omitempty"`
	// 身份证正面
	LegalCertBackUrl string `json:"legal_cert_back_url,omitempty" xml:"legal_cert_back_url,omitempty"`
	// 身份证反面
	LegalCertFrontUrl string `json:"legal_cert_front_url,omitempty" xml:"legal_cert_front_url,omitempty"`
	// 身份证有效期
	LegalIdExpire string `json:"legal_id_expire,omitempty" xml:"legal_id_expire,omitempty"`
	// 公司注册地址
	CompanyAddress string `json:"company_address,omitempty" xml:"company_address,omitempty"`
	// 营业执照地址
	BusinessLicenseUrl string `json:"business_license_url,omitempty" xml:"business_license_url,omitempty"`
	// 是否使用pos
	UsePos int64 `json:"use_pos,omitempty" xml:"use_pos,omitempty"`
}
