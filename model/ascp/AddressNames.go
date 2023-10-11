package ascp

// AddressNames 结构体
type AddressNames struct {
	// 浙江省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}
