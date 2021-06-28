package tmallservice

// TmallServiceSettleadjustmentCancelResult 
/* model for simplify = false
type TmallServiceSettleadjustmentCancelResult struct {

    // errorMessage
    
    ErrorMessage  *struct {
        ErrorMessage  *ErrorMessage `json:"error_message,omitempty"`
    } `json:"error_message,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TmallServiceSettleadjustmentCancelResult 
type TmallServiceSettleadjustmentCancelResult struct {

    // errorMessage
    ErrorMessage   *ErrorMessage `json:"error_message,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
