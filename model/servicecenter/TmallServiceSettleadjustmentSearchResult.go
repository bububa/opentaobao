package servicecenter

// TmallServiceSettleadjustmentSearchResult 
/* model for simplify = false
type TmallServiceSettleadjustmentSearchResult struct {

    // dataModule
    
    SettleAdjustmentList  struct {
        SettleAdjustmentResponse  []SettleAdjustmentResponse `json:"settle_adjustment_response,omitempty"`
    } `json:"settle_adjustment_list,omitempty"`
    

    // errorMessage
    
    ErrorMessage  *struct {
        ErrorMessage  *ErrorMessage `json:"error_message,omitempty"`
    } `json:"error_message,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TmallServiceSettleadjustmentSearchResult 
type TmallServiceSettleadjustmentSearchResult struct {

    // dataModule
    SettleAdjustmentList   []SettleAdjustmentResponse `json:"settle_adjustment_list,omitempty"`

    // errorMessage
    ErrorMessage   *ErrorMessage `json:"error_message,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
