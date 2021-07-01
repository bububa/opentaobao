package nrt

// TmallNrtPayMerchantStallSigningModifyResultDo 结构体
type TmallNrtPayMerchantStallSigningModifyResultDo struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 系统自动生成
	Data *StallSigningRespDto `json:"data,omitempty" xml:"data,omitempty"`
	// 系统自动生成
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 系统自动生成
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
}
