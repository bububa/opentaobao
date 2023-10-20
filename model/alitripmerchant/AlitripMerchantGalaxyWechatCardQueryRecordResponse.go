package alitripmerchant

// AlitripmerchantgalaxywechatcardqueryrecordResponse 结构体
type AlitripmerchantgalaxywechatcardqueryrecordResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
