package tmallservice

// AddressDto 结构体
type AddressDto struct {
	// 省/市/区/街道
	AddressNames []string `json:"address_names,omitempty" xml:"address_names>string,omitempty"`
	// 详细地址，街到门牌，
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
}
