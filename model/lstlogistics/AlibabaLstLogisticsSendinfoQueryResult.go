package lstlogistics

// AlibabaLstLogisticsSendinfoQueryResult 结构体
type AlibabaLstLogisticsSendinfoQueryResult struct {
	// 返回内容
	ContentList []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
