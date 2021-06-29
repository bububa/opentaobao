package icbudropshipping

// Inventory 
type Inventory struct {
    // inventory count
    Inventory   int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
    // store code
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    // dispatch country
    DispatchCountry   string `json:"dispatch_country,omitempty" xml:"dispatch_country,omitempty"`
}
