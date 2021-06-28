package tmallservice

// TmallServiceSettleadjustmentGetResult 
/* model for simplify = false
type TmallServiceSettleadjustmentGetResult struct {

    // dataModule
    
    DataModule  *struct {
        SettleAdjustmentResponse  *SettleAdjustmentResponse `json:"settle_adjustment_response,omitempty"`
    } `json:"data_module,omitempty"`
    

    // errorMessage
    
    ErrorMessage  *struct {
        ErrorMessage  *ErrorMessage `json:"error_message,omitempty"`
    } `json:"error_message,omitempty"`
    

    // success，成功true，失败false
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TmallServiceSettleadjustmentGetResult 
type TmallServiceSettleadjustmentGetResult struct {

    // dataModule
    DataModule   *SettleAdjustmentResponse `json:"data_module,omitempty"`

    // errorMessage
    ErrorMessage   *ErrorMessage `json:"error_message,omitempty"`

    // success，成功true，失败false
    Success   bool `json:"success,omitempty"`

}
