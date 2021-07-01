package omniorder

// TaobaoOmniorderStoreSdtcancelResult 结构体
type TaobaoOmniorderStoreSdtcancelResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
}
