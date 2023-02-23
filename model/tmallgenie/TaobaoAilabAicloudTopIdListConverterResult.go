package tmallgenie

// TaobaoAilabAicloudTopIdListConverterResult 结构体
type TaobaoAilabAicloudTopIdListConverterResult struct {
	// 返回查询内容
	RtValue []OpenInfoResponse `json:"rt_value,omitempty" xml:"rt_value>open_info_response,omitempty"`
	// 返回错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回code
	RtCode int64 `json:"rt_code,omitempty" xml:"rt_code,omitempty"`
	// 请求状态
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
