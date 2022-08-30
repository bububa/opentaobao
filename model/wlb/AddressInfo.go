package wlb

// AddressInfo 结构体
type AddressInfo struct {
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Borough string `json:"borough,omitempty" xml:"borough,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 邮编
	Zip string `json:"zip,omitempty" xml:"zip,omitempty"`
}
