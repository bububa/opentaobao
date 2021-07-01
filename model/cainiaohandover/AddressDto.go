package cainiaohandover

// AddressDto 结构体
type AddressDto struct {
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 街道
	Street string `json:"street,omitempty" xml:"street,omitempty"`
	// 区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
}
