package pur

// AlibabaCeresSupplierPoQueryResult 
type AlibabaCeresSupplierPoQueryResult struct {

    // 返回信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 返回单据消息体List
    
    Values   []Value `json:"values,omitempty" xml:"values,omitempty"`
    

    // 返回code
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 是否查询成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
