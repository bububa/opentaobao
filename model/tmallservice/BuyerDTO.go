package tmallservice

// BuyerDTO 
type BuyerDTO struct {

    // 省
    
    AddressProvince   string `json:"address_province,omitempty" xml:"address_province,omitempty"`
    

    // 详细地址
    
    AddressDetail   string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
    

    // 买家淘宝账号
    
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    

    // 买家电话
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 买家手机号
    
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    

    // 地区编码
    
    Location   int64 `json:"location,omitempty" xml:"location,omitempty"`
    

    // 区
    
    AddressDistrict   string `json:"address_district,omitempty" xml:"address_district,omitempty"`
    

    // 买家姓名
    
    BuyerName   string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
    

    // 街道
    
    AddressTown   string `json:"address_town,omitempty" xml:"address_town,omitempty"`
    

    // 市
    
    AddressCity   string `json:"address_city,omitempty" xml:"address_city,omitempty"`
    

}
