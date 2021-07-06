package omniorder

// TaobaoOmniorderStoreGetconsignmailcodeResult 结构体
type TaobaoOmniorderStoreGetconsignmailcodeResult struct {
	// code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 发货信息
	Data *GetStoreConsignCodeResponse `json:"data,omitempty" xml:"data,omitempty"`
}
