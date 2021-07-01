package lstlogistics

// AlibabaLstLogisticsTraceQueryResult 结构体
type AlibabaLstLogisticsTraceQueryResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回内容
	ContentList []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
}
