package tmallservice

// TmallServiceSettleadjustmentRequestResult 
/* model for simplify = false
type TmallServiceSettleadjustmentRequestResult struct {

    // dataModule
    
    DataModule  *struct {
        SettleAdjustmentResp  *SettleAdjustmentResp `json:"settle_adjustment_resp,omitempty"`
    } `json:"data_module,omitempty"`
    

    // errorMessage
    
    ErrorMessage  *struct {
        ErrorMessage  *ErrorMessage `json:"error_message,omitempty"`
    } `json:"error_message,omitempty"`
    

    // true：查询成功，false：失败
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TmallServiceSettleadjustmentRequestResult 
type TmallServiceSettleadjustmentRequestResult struct {

    // dataModule
    DataModule   *SettleAdjustmentResp `json:"data_module,omitempty"`

    // errorMessage
    ErrorMessage   *ErrorMessage `json:"error_message,omitempty"`

    // true：查询成功，false：失败
    Success   bool `json:"success,omitempty"`

}
