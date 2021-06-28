package mos

// DeliveryCustomDto 
/* model for simplify = false
type DeliveryCustomDto struct {

    // 名
    
    Name   string `json:"name,omitempty"`
    

    // 头像
    
    AvatarUrl   string `json:"avatar_url,omitempty"`
    

    // 电话
    
    MobilePhone   string `json:"mobile_phone,omitempty"`
    

    // 电话
    
    Telphone   string `json:"telphone,omitempty"`
    

    // 详细信息
    
    AddressInfo  *struct {
        DeliveryAddressDto  *DeliveryAddressDto `json:"delivery_address_dto,omitempty"`
    } `json:"address_info,omitempty"`
    

}
*/

// DeliveryCustomDto 
type DeliveryCustomDto struct {

    // 名
    Name   string `json:"name,omitempty"`

    // 头像
    AvatarUrl   string `json:"avatar_url,omitempty"`

    // 电话
    MobilePhone   string `json:"mobile_phone,omitempty"`

    // 电话
    Telphone   string `json:"telphone,omitempty"`

    // 详细信息
    AddressInfo   *DeliveryAddressDto `json:"address_info,omitempty"`

}
