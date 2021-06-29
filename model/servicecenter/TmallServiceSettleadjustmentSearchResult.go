package servicecenter

// TmallServiceSettleadjustmentSearchResult 
type TmallServiceSettleadjustmentSearchResult struct {
    // dataModule
    SettleAdjustmentList   []SettleAdjustmentResponse `json:"settle_adjustment_list,omitempty" xml:"settle_adjustment_list>settle_adjustment_response,omitempty"`
    // errorMessage
    ErrorMessage   *ErrorMessage `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
