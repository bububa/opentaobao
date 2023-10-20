package tbk

// RegisterInfoDto 结构体
type RegisterInfoDto struct {
	// 渠道独有 -店铺名称
	Shopname string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 渠道独有 -店铺类型
	Shoptype string `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
	// 渠道独有 - 信息
	Phonenumber string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 渠道独有 -信息
	Detailaddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 渠道独有 -区
	Location string `json:"location,omitempty" xml:"location,omitempty"`
	// 渠道独有 -类型
	Shopcertifytype string `json:"shop_certify_type,omitempty" xml:"shop_certify_type,omitempty"`
	// 渠道独有 - 编号
	Certifynumber string `json:"certify_number,omitempty" xml:"certify_number,omitempty"`
	// 渠道独有 -经营类型
	Career string `json:"career,omitempty" xml:"career,omitempty"`
}
