package tmallservice

// TmallServiceSettleadjustmentRequestResult 
type TmallServiceSettleadjustmentRequestResult struct {
    // dataModule
    DataModule   *SettleAdjustmentResp `json:"data_module,omitempty" xml:"data_module,omitempty"`
    // errorMessage
    ErrorMessage   *ErrorMessage `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // true：查询成功，false：失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
