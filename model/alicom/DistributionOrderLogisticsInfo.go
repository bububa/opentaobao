package alicom

// DistributionOrderLogisticsInfo 
type DistributionOrderLogisticsInfo struct {

    // 区
    
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    

    // 市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 国家
    
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    

    // 省
    
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    

    // 详细地址
    
    RawAddress   string `json:"raw_address,omitempty" xml:"raw_address,omitempty"`
    

    // 乡/镇/街道
    
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    

    // 收货人名称
    
    ConsigneeName   string `json:"consignee_name,omitempty" xml:"consignee_name,omitempty"`
    

    // 收货人联系电话
    
    ContactPhone   string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
    

    // 收货人地址
    
    DeliveryAddress   string `json:"delivery_address,omitempty" xml:"delivery_address,omitempty"`
    

}
