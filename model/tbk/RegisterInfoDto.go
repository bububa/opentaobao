package tbk

// RegisterInfoDto 结构体
type RegisterInfoDto struct {
	// 渠道独有 -店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 渠道独有 -店铺类型
	ShopType string `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
	// 渠道独有 -手机号码
	PhoneNumber string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 渠道独有 -详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 渠道独有 -地区
	Location string `json:"location,omitempty" xml:"location,omitempty"`
	// 渠道独有 -证件类型
	ShopCertifyType string `json:"shop_certify_type,omitempty" xml:"shop_certify_type,omitempty"`
	// 渠道独有 -对应的证件证件类型编号
	CertifyNumber string `json:"certify_number,omitempty" xml:"certify_number,omitempty"`
	// 渠道独有 -经营类型
	Career string `json:"career,omitempty" xml:"career,omitempty"`
}
