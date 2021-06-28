package tmallservice

// AddressDto 
/* model for simplify = false
type AddressDto struct {

    // 详细地址，街到门牌，
    
    AddressDetail   string `json:"address_detail,omitempty"`
    

    // 省/市/区/街道
    
    AddressNames  struct {
        String  []string `json:"string,omitempty"`
    } `json:"address_names,omitempty"`
    

}
*/

// AddressDto 
type AddressDto struct {

    // 详细地址，街到门牌，
    AddressDetail   string `json:"address_detail,omitempty"`

    // 省/市/区/街道
    AddressNames   []string `json:"address_names,omitempty"`

}
