package logistic

// JzReceiverTo 结构体
type JzReceiverTo struct {
	// 座机号
	TelePhone string `json:"tele_phone,omitempty" xml:"tele_phone,omitempty"`
	// 手机号
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 收货人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 街道
	Street string `json:"street,omitempty" xml:"street,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
}
