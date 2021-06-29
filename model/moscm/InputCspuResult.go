package moscm

// InputCspuResult 
type InputCspuResult struct {

    // 中台商品id
    
    CspuId   string `json:"cspu_id,omitempty" xml:"cspu_id,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 供应商商品id
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // 是否录入成功，true：成功 false：失败
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
