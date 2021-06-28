package wdk

// AlibabaPosFundCashierShiftSummaryResult 
/* model for simplify = false
type AlibabaPosFundCashierShiftSummaryResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 扩展字段
    
    BizExtMap   string `json:"biz_ext_map,omitempty"`
    

    // 模型
    
    Model  struct {
        CashierShiftFundSummaryDTO  []CashierShiftFundSummaryDTO `json:"cashier_shift_fund_summary_dto,omitempty"`
    } `json:"model,omitempty"`
    

}
*/

// AlibabaPosFundCashierShiftSummaryResult 
type AlibabaPosFundCashierShiftSummaryResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // 扩展字段
    BizExtMap   string `json:"biz_ext_map,omitempty"`

    // 模型
    Model   []CashierShiftFundSummaryDTO `json:"model,omitempty"`

}
