package user

// TmallServiceSettleadjustmentModifyResult 
/* model for simplify = false
type TmallServiceSettleadjustmentModifyResult struct {

    // errorMessage
    
    ErrorMessage  *struct {
        ErrorMessage  *ErrorMessage `json:"error_message,omitempty"`
    } `json:"error_message,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TmallServiceSettleadjustmentModifyResult 
type TmallServiceSettleadjustmentModifyResult struct {

    // errorMessage
    ErrorMessage   *ErrorMessage `json:"error_message,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
