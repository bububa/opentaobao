package tmallservice

// TmallServiceSettleadjustmentGetResult 
type TmallServiceSettleadjustmentGetResult struct {

    // dataModule
    DataModule   *SettleAdjustmentResponse `json:"data_module,omitempty"`

    // errorMessage
    ErrorMessage   *ErrorMessage `json:"error_message,omitempty"`

    // success，成功true，失败false
    Success   bool `json:"success,omitempty"`

}
