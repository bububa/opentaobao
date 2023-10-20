package car

// TaobaoalitripcarorderacceptApiResult 结构体
type TaobaoalitripcarorderacceptApiResult struct {
	// 其它数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码 0成功 其它见文档
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 错误码 0成功 其它见文档
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
}
