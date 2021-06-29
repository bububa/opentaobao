package icbudropshipping

// AddressInfoDTO 
type AddressInfoDTO struct {
    // Shipping address
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // City
    City   *DivisionInfoDTO `json:"city,omitempty" xml:"city,omitempty"`
    // Country
    Country   *DivisionInfoDTO `json:"country,omitempty" xml:"country,omitempty"`
    // province
    Province   *DivisionInfoDTO `json:"province,omitempty" xml:"province,omitempty"`
    // If any, please send it to us to make the freight more accurate.
    Zip   string `json:"zip,omitempty" xml:"zip,omitempty"`
}
