package moscm

// PriceResult 
type PriceResult struct {

    // 如果出错，返回错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 价格对象
    
    PriceDto   *PriceDto `json:"price_dto,omitempty" xml:"price_dto,omitempty"`
    

    // 单挑成功与否
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
