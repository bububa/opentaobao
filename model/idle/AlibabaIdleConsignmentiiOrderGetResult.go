package idle

// AlibabaIdleConsignmentiiOrderGetResult 
type AlibabaIdleConsignmentiiOrderGetResult struct {

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 订单详情
    
    Module   *ConsignmentV2OrderTo `json:"module,omitempty" xml:"module,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
