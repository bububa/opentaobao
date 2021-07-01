package wenyuvideo

// YoukuWenyuvideoPersionGetResult 结构体
type YoukuWenyuvideoPersionGetResult struct {
	// 人物详情信息，包含相关影视和相关人物
	Model *PersonDetailRbo `json:"model,omitempty" xml:"model,omitempty"`
	// HTTP请求状态
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// 业务扩展数据
	BizExtMap string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// 错误码，业务错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 业务错误提示
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
