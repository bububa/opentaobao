package cainiaolocker

// AddressDto 结构体
type AddressDto struct {
	// 城市，长度小于20
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 详细地址，长度小于256
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 区，长度小于20
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 省，长度小于20
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 街道，长度小于30
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}
