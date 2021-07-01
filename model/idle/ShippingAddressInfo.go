package idle

// ShippingAddressInfo 结构体
type ShippingAddressInfo struct {
	// 这里必须写全详细地址
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收件人姓名
	ConsigneeName string `json:"consignee_name,omitempty" xml:"consignee_name,omitempty"`
	// 手机号码
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 邮编
	PostCode string `json:"post_code,omitempty" xml:"post_code,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 固话
	TelNo string `json:"tel_no,omitempty" xml:"tel_no,omitempty"`
	// 四级地址 镇/街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}
