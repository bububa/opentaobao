package logistic

// ShipFresh 
type ShipFresh struct {

    // 物流公司名称
    
    DeliveryName   string `json:"delivery_name,omitempty" xml:"delivery_name,omitempty"`
    

    // 物流公司服务电话
    
    Tel   string `json:"tel,omitempty" xml:"tel,omitempty"`
    

    // 旺旺ID
    
    WangwangId   string `json:"wangwang_id,omitempty" xml:"wangwang_id,omitempty"`
    

    // 返回的地址信息
    
    SendAddr   string `json:"send_addr,omitempty" xml:"send_addr,omitempty"`
    

}
