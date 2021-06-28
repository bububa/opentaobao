package bus

// TaobaoAlitripBusTicketsInsuranceRecommendResult 
/* model for simplify = false
type TaobaoAlitripBusTicketsInsuranceRecommendResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 接口返回结果对象
    
    Response  *struct {
        TopStandardInsRecommendResponse  *TopStandardInsRecommendResponse `json:"top_standard_ins_recommend_response,omitempty"`
    } `json:"response,omitempty"`
    

    // 扩展预留
    
    BizExtMap   string `json:"biz_ext_map,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

}
*/

// TaobaoAlitripBusTicketsInsuranceRecommendResult 
type TaobaoAlitripBusTicketsInsuranceRecommendResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 接口返回结果对象
    Response   *TopStandardInsRecommendResponse `json:"response,omitempty"`

    // 扩展预留
    BizExtMap   string `json:"biz_ext_map,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

}
