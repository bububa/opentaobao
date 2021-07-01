package user

// TmallServiceSettleadjustmentModifyResult 结构体
type TmallServiceSettleadjustmentModifyResult struct {
	// errorMessage
	ErrorMessage *ErrorMessage `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
