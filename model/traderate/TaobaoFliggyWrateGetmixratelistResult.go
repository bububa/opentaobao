package traderate

// TaobaoFliggyWrateGetmixratelistResult 
/* model for simplify = false
type TaobaoFliggyWrateGetmixratelistResult struct {

    // 返回对象
    
    Model  *struct {
        GetMixRateListResult  *GetMixRateListResult `json:"get_mix_rate_list_result,omitempty"`
    } `json:"model,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoFliggyWrateGetmixratelistResult 
type TaobaoFliggyWrateGetmixratelistResult struct {

    // 返回对象
    Model   *GetMixRateListResult `json:"model,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
