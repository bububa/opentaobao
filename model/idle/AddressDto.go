package idle

// AddressDTO 
type AddressDTO struct {
    // 区
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 省
    Prov   string `json:"prov,omitempty" xml:"prov,omitempty"`
}
