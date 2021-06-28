package servicecenter

// ResultVo 
type ResultVo struct {

    // 消耗时间
    
    CostTime   int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
    

    // 当前时间
    
    GmtCurrentTime   int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
    

    // 异常代码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

    // 异常消息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // 真正返回对象
    
    Object   *LeaseOrderInfoDto `json:"object,omitempty" xml:"object,omitempty"`
    

    // 成功与否
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
