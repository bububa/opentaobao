package idle

// AlibabaIdleRecycleOrderShowResult 
type AlibabaIdleRecycleOrderShowResult struct {

    // errMsg
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 订单信息
    
    Module   *Serializable `json:"module,omitempty" xml:"module,omitempty"`
    

}
