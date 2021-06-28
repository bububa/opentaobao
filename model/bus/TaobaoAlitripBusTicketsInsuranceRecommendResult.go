package bus

// TaobaoAlitripBusTicketsInsuranceRecommendResult 
type TaobaoAlitripBusTicketsInsuranceRecommendResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 接口返回结果对象
    
    Response   *TopStandardInsRecommendResponse `json:"response,omitempty" xml:"response,omitempty"`
    

    // 扩展预留
    
    BizExtMap   string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

}
