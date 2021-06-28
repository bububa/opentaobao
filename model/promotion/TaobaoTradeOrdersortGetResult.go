package promotion

// TaobaoTradeOrdersortGetResult 
/* model for simplify = false
type TaobaoTradeOrdersortGetResult struct {

    // 返回素材id
    
    Data  *struct {
        AwardOrderResultDto  *AwardOrderResultDto `json:"award_order_result_dto,omitempty"`
    } `json:"data,omitempty"`
    

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoTradeOrdersortGetResult 
type TaobaoTradeOrdersortGetResult struct {

    // 返回素材id
    Data   *AwardOrderResultDto `json:"data,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
