package promotion

// TaobaoCardExpandcardQueryResult 结构体
type TaobaoCardExpandcardQueryResult struct {
	// 0为成功，其他为失败
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 卡信息
	Models []ExpandCardVo `json:"models,omitempty" xml:"models>expand_card_vo,omitempty"`
	// debugInfo
	DebugInfo string `json:"debug_info,omitempty" xml:"debug_info,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误级别
	ErrorLevel string `json:"error_level,omitempty" xml:"error_level,omitempty"`
}
