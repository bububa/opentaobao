package ascp

// OrderReceiverPrivacyData 结构体
type OrderReceiverPrivacyData struct {
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 移动电话隐私
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 国家二字码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 收件人所在城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收件人所在区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 收件人所在街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}
