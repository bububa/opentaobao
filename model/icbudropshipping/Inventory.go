package icbudropshipping

// Inventory 结构体
type Inventory struct {
	// dispatch country
	DispatchCountry string `json:"dispatch_country,omitempty" xml:"dispatch_country,omitempty"`
	// store code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// inventory count
	Inventory int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
}
