package axintrade

// TaobaoalitripaxintransfundconfirmResult 结构体
type TaobaoalitripaxintransfundconfirmResult struct {
	// 描述信息
	InfoMsg string `json:"info_msg,omitempty" xml:"info_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 接口出参
	Data *AxinFundConfirmResDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否需要重试
	NeedRetry bool `json:"need_retry,omitempty" xml:"need_retry,omitempty"`
}
