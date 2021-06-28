package car

// OrderComplete 
type OrderComplete struct {

    // 服务商生成的订单ID
    
    ThirdOrderId   string `json:"third_order_id,omitempty" xml:"third_order_id,omitempty"`
    

    // 服务商标识
    
    ProviderId   string `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
    

    // 服务完成时间
    
    CompleteTime   string `json:"complete_time,omitempty" xml:"complete_time,omitempty"`
    

    // 阿里旅行生成的订单ID
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 可选，卖家id
    
    SellerId   string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    

    // 0:接送机 1：实时打车 2：租车(不传值默认为0)
    
    UseType   int64 `json:"use_type,omitempty" xml:"use_type,omitempty"`
    

    // 价格详情
    
    PriceInfo   *PriceInfo `json:"price_info,omitempty" xml:"price_info,omitempty"`
    

    // 实际行驶公里数
    
    Distance   string `json:"distance,omitempty" xml:"distance,omitempty"`
    

}
