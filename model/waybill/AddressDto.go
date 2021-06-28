package waybill

// AddressDto 
/* model for simplify = false
type AddressDto struct {

    // 城市，长度小于20
    
    City   string `json:"city,omitempty"`
    

    // 详细地址，长度小于256
    
    Detail   string `json:"detail,omitempty"`
    

    // 区，长度小于20
    
    District   string `json:"district,omitempty"`
    

    // 省，长度小于20
    
    Province   string `json:"province,omitempty"`
    

    // 街道，长度小于30
    
    Town   string `json:"town,omitempty"`
    

}
*/

// AddressDto 
type AddressDto struct {

    // 城市，长度小于20
    City   string `json:"city,omitempty"`

    // 详细地址，长度小于256
    Detail   string `json:"detail,omitempty"`

    // 区，长度小于20
    District   string `json:"district,omitempty"`

    // 省，长度小于20
    Province   string `json:"province,omitempty"`

    // 街道，长度小于30
    Town   string `json:"town,omitempty"`

}
