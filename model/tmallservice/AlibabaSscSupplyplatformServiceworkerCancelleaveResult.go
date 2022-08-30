package tmallservice

// AlibabaSscSupplyplatformServiceworkerCancelleaveResult 结构体
type AlibabaSscSupplyplatformServiceworkerCancelleaveResult struct {
	// 错误码
	DisplayCode string `json:"display_code,omitempty" xml:"display_code,omitempty"`
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 是否成功
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
