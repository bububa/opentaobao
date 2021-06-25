package servicecenter

// TmallServiceSettleadjustmentSearchResult 
type TmallServiceSettleadjustmentSearchResult struct {

    // dataModule
    SettleAdjustmentList   []SettleAdjustmentResponse `json:"settle_adjustment_list,omitempty"`

    // errorMessage
    ErrorMessage   *ErrorMessage `json:"error_message,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
