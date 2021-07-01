package idle

// UserAddressDto 结构体
type UserAddressDto struct {
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 发货人名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 省
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 手机号
	PhoneNo string `json:"phone_no,omitempty" xml:"phone_no,omitempty"`
}
