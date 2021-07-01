package icbudropshipping

// AddressInfoDto 结构体
type AddressInfoDto struct {
	// Shipping address
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// City
	City *DivisionInfoDto `json:"city,omitempty" xml:"city,omitempty"`
	// Country
	Country *DivisionInfoDto `json:"country,omitempty" xml:"country,omitempty"`
	// province
	Province *DivisionInfoDto `json:"province,omitempty" xml:"province,omitempty"`
	// If any, please send it to us to make the freight more accurate.
	Zip string `json:"zip,omitempty" xml:"zip,omitempty"`
}
