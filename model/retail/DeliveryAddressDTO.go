package retail

// DeliveryAddressDto 结构体
type DeliveryAddressDto struct {
	// 收货人姓名
	FullName string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 街道地址
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}
