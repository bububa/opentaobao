package omniorder

// ReceiverDto 
type ReceiverDto struct {

    // 收件人详细地址
    
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    

    // 城市
    
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    

    // 区县
    
    DistrictName   string `json:"district_name,omitempty" xml:"district_name,omitempty"`
    

    // 收件人手机号，该字段与收件人电话二者必填至少其一
    
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    

    // 收件人电话
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 省
    
    ProName   string `json:"pro_name,omitempty" xml:"pro_name,omitempty"`
    

    // 收件人姓名
    
    ReceiverName   string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
    

    // 街道
    
    StreetName   string `json:"street_name,omitempty" xml:"street_name,omitempty"`
    

}
